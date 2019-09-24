package order

import (
	"context"
	"sync"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	productSvc "github.com/phungvandat/life-cafe-backend/service/product"
	userSvc "github.com/phungvandat/life-cafe-backend/service/user"
	"github.com/phungvandat/life-cafe-backend/util/contextkey"
	errors "github.com/phungvandat/life-cafe-backend/util/error"
	"github.com/phungvandat/life-cafe-backend/util/externals/sagas"
	"github.com/phungvandat/life-cafe-backend/util/helper"
)

// pgService implmenter for auth service in postgres
type pgService struct {
	db         *gorm.DB
	userSvc    userSvc.Service
	productSvc productSvc.Service
	spRollback sagas.SagasService
}

// NewPGService new pg service
func NewPGService(db *gorm.DB, userSvc userSvc.Service, productSvc productSvc.Service, spRollback sagas.SagasService) Service {
	return &pgService{
		db:         db,
		spRollback: spRollback,
		userSvc:    userSvc,
		productSvc: productSvc,
	}
}

func (s *pgService) RollbackTransaction(_ context.Context, transactionID string) error {
	return s.spRollback.RollbackTransaction(transactionID)
}

func (s *pgService) CommitTransaction(_ context.Context, transactionID string) error {
	return s.spRollback.CommitTransaction(transactionID)
}

func (s *pgService) rollbackTransactions(ctx context.Context, transactionIDs []string) {
	var log log.Logger
	for _, transactionID := range transactionIDs {
		err := s.RollbackTransaction(ctx, transactionID)
		if err != nil {
			log.Log("Rollback transaction "+transactionID+" failure by error ", err)
		}
	}
}

func (s *pgService) commitTransactions(ctx context.Context, transactionIDs []string) {
	var log log.Logger
	for _, transactionID := range transactionIDs {
		err := s.CommitTransaction(ctx, transactionID)
		if err != nil {
			log.Log("Commit transaction "+transactionID+" failure by error ", err)
		}
	}
}

func (s *pgService) CreateOrder(ctx context.Context, req requestModel.CreateOrderRequest) (*responseModel.CreateOrderResponse, error) {
	tx := s.db.Begin()
	transactionID := (pgModel.NewUUID()).String()
	s.spRollback.NewTransaction(transactionID, tx)
	res := &responseModel.CreateOrderResponse{
		Order:         &responseModel.Order{},
		TransactionID: &transactionID,
	}

	// arrOtherServiceTX is array other service transaction commit or rollback
	arrOtherServiceTX := []string{}

	ctxUserID, check := ctx.Value(contextkey.UserIDContextKey).(string)
	if !check {
		return res, errors.NotLoggedInError
	}
	signUserID, _ := pgModel.UUIDFromString(ctxUserID)

	var err error
	order := &pgModel.Order{
		Type:                req.Type,
		CreatorID:           &signUserID,
		Note:                req.Note,
		Status:              req.Status,
		ReceiverPhoneNumber: req.ReceiverPhoneNumber,
		ReceiverAddress:     req.ReceiverAddress,
		ReceiverFullname:    req.ReceiverFullname,
	}

	if order.Status == "done" {
		order.ImplementerID = &signUserID
	}

	var customer *pgModel.User
	// Get order customer
	if req.Type != "import" {
		if req.CustomerID == "" {
			createCustomerReq := requestModel.CreateUserRequest{
				Username:    req.CustomerPhoneNumber,
				PhoneNumber: req.CustomerPhoneNumber,
				Fullname:    req.CustomerFullname,
				Address:     req.CustomerAddress,
				Role:        "user",
			}

			customerRes, err := s.userSvc.Create(ctx, createCustomerReq)
			if customerRes != nil && customerRes.TransactionID != nil {
				arrOtherServiceTX = append(arrOtherServiceTX, *customerRes.TransactionID)
			}

			if err != nil {
				s.rollbackTransactions(ctx, arrOtherServiceTX)
				return res, err
			}
			customer = customerRes.User
		} else {
			getCustomerReq := requestModel.GetUserRequest{
				ParamUserID: req.CustomerID,
			}

			customerRes, err := s.userSvc.GetUser(ctx, getCustomerReq)
			if err != nil {
				s.rollbackTransactions(ctx, arrOtherServiceTX)
				return res, err
			}

			customer = customerRes.User
		}

		order.CustomerID = &customer.ID
	}
	err = tx.Create(order).Error

	if err != nil {
		return res, err
	}
	orderProductInfo := []*responseModel.OrderProductInfo{}

	// Order product info
	wg := sync.WaitGroup{}

	for _, productInfo := range req.OrderProductInfo {
		wg.Add(1)
		go func(productInfo requestModel.ProductOrder) {
			defer wg.Done()
			productRes, funcErr := s.productSvc.GetProduct(ctx, requestModel.GetProductRequest{
				ParamProductID: productInfo.ProductID,
			})
			if funcErr != nil {
				err = funcErr
				return
			}

			product := productRes.Product
			if order.Type == "export" {
				if productInfo.OrderQuantity > product.Quantity {
					err = errors.InvalidOrderProductQuantityError
					return
				}
				product.Quantity -= productInfo.OrderQuantity
			} else if order.Type == "import" && order.Status == "done" {
				product.Quantity += productInfo.OrderQuantity
			}

			productOrder := &pgModel.ProductOrder{
				ProductID:      &product.ID,
				OrderID:        &order.ID,
				OrderQuantity:  productInfo.OrderQuantity,
				OrderRealPrice: productInfo.OrderRealPrice,
			}

			funcErr = tx.Create(productOrder).Error
			if funcErr != nil {
				err = funcErr
				return
			}
			updateProductRes, funcErr := s.productSvc.UpdateProduct(ctx,
				requestModel.UpdateProductRequest{
					ParamProductID: (product.ID).String(),
					Quantity:       &product.Quantity,
				})
			if updateProductRes != nil && updateProductRes.TransactionID != nil {
				arrOtherServiceTX = append(arrOtherServiceTX, *updateProductRes.TransactionID)
			}
			if funcErr != nil {
				err = funcErr
				return
			}
			orderProductInfo = append(orderProductInfo, &responseModel.OrderProductInfo{
				OrderQuantity:  productOrder.OrderQuantity,
				OrderRealPrice: productOrder.OrderRealPrice,
				Product: &pgModel.Product{
					Model: pgModel.Model{
						ID: product.ID,
					},
					Name:  product.Name,
					Color: product.Color,
				},
			})
		}(productInfo)
	}
	wg.Wait()

	if err != nil {
		s.rollbackTransactions(ctx, arrOtherServiceTX)
		return res, err
	}

	res.Order.Order = order
	if customer != nil {
		res.Order.Customer = &pgModel.User{
			Model: pgModel.Model{
				ID: customer.ID,
			},
			Fullname:    customer.Fullname,
			Address:     customer.Address,
			PhoneNumber: customer.PhoneNumber,
		}
	}
	res.Order.OrderProductInfo = orderProductInfo
	s.commitTransactions(ctx, arrOtherServiceTX)

	return res, nil
}

func (s *pgService) GetOrder(ctx context.Context, req requestModel.GetOrderRequest) (*responseModel.GetOrderResponse, error) {
	res := &responseModel.GetOrderResponse{
		Order: &responseModel.Order{},
	}

	orderID, _ := pgModel.UUIDFromString(req.ParamOrderID)

	order := &pgModel.Order{
		Model: pgModel.Model{
			ID: orderID,
		},
	}

	err := s.db.Find(order, order).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.OrderNotExistError
		}
		return res, err
	}

	// check permission
	ctxUserID, _ := ctx.Value(contextkey.UserIDContextKey).(string)

	userSiginRole := ctx.Value(contextkey.UserRoleContextKey).(string)

	if userSiginRole != "amdin" && userSiginRole != "master" && order.CustomerID != nil && ctxUserID != (order.CustomerID).String() && ctxUserID != (order.CreatorID).String() && order.ImplementerID != nil && (order.ImplementerID).String() != ctxUserID {
		return res, errors.PermissionDeniedError
	}

	orderRes, err := s.getOrderDetail(ctx, *order, true)

	if err != nil {
		return res, err
	}

	res.Order = orderRes

	return res, nil
}

func (s *pgService) GetOrders(ctx context.Context, req requestModel.GetOrdersRequest) (*responseModel.GetOrdersResponse, error) {
	res := &responseModel.GetOrdersResponse{}

	skip := req.Skip
	limit := req.Limit

	if req.Skip == "" {
		skip = "-1"
	}

	if req.Limit == "" {
		limit = "-1"
	}

	var total int
	var err error
	orders := []pgModel.Order{}

	findOrdersWG := sync.WaitGroup{}

	findOrdersWG.Add(1)
	go func() {
		defer findOrdersWG.Done()
		errFunc := s.db.Offset(skip).Limit(limit).Order("created_at desc").Find(&orders).Error
		if errFunc != nil {
			err = errFunc
		}
	}()

	findOrdersWG.Add(1)
	go func() {
		defer findOrdersWG.Done()
		errFunc := s.db.Table("orders").Count(&total).Error
		if errFunc != nil {
			err = errFunc
		}
	}()

	findOrdersWG.Wait()

	if err != nil {
		return res, err
	}

	orderArr := make([]*responseModel.Order, len(orders))

	// Get order detail
	wg := sync.WaitGroup{}

	for idx, order := range orders {
		wg.Add(1)
		go func(order pgModel.Order, idx int) {
			defer wg.Done()
			orderRes, errFunc := s.getOrderDetail(ctx, order, true)

			if errFunc != nil {
				err = errFunc
				return
			}
			orderArr[idx] = orderRes
		}(order, idx)
	}

	wg.Wait()

	if err != nil {
		return res, err
	}

	res.Orders = orderArr
	res.Total = total

	return res, nil
}

func (s *pgService) UpdateOrder(ctx context.Context, req requestModel.UpdateOrderRequest) (*responseModel.UpdateOrderResponse, error) {
	tx := s.db.Begin()
	transactionID := (pgModel.NewUUID()).String()
	s.spRollback.NewTransaction(transactionID, tx)
	res := &responseModel.UpdateOrderResponse{
		TransactionID: &transactionID,
	}

	arrOtherServiceTX := []string{}

	orderIDUUID, _ := pgModel.UUIDFromString(req.ParamOrderID)
	orderProductInfo := []*responseModel.OrderProductInfo{}
	wg := sync.WaitGroup{}

	order := &pgModel.Order{
		Model: pgModel.Model{
			ID: orderIDUUID,
		},
	}

	err := s.db.Find(order, order).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.OrderNotExistError
		}
		s.rollbackTransactions(ctx, arrOtherServiceTX)
		return res, err
	}

	if (order.Status == "delivering" && req.Status != "done") || order.Status == "done" {
		s.rollbackTransactions(ctx, arrOtherServiceTX)
		return res, errors.CannotUpdateOrderError
	}

	if req.Note != "" && req.Note != order.Note {
		order.Note = req.Note
	}

	if req.Status != "" && req.Status != order.Status {
		preStatus := order.Status
		order.Status = req.Status
		if req.Status == "done" {
			ctxUserID := ctx.Value(contextkey.UserIDContextKey).(string)
			implementerID, _ := pgModel.UUIDFromString(ctxUserID)
			order.ImplementerID = &implementerID
		}
		if preStatus == "delivering" {
			goto Update
		}
	}

	if len(req.OrderProductInfo) > 0 {
		preProductOrders := []pgModel.ProductOrder{}
		err = s.db.Where("order_id = ?", (order.ID).String()).Find(&preProductOrders).Error

		if err != nil {
			s.rollbackTransactions(ctx, arrOtherServiceTX)
			return res, err
		}
		preProductIDs := []string{}

		for _, value := range preProductOrders {
			preProductIDs = append(preProductIDs, (value.ProductID).String())
		}
		currentProductIDs := []string{}

		for _, value := range req.OrderProductInfo {
			currentProductIDs = append(currentProductIDs, value.ProductID)
		}

		sameProductIDs := helper.GetSameElementInArrays(preProductIDs, currentProductIDs)

		deleteProductIDs := helper.DifferenceArray(preProductIDs, sameProductIDs)

		// Delete product order
		for _, productID := range deleteProductIDs {
			wg.Add(1)
			go func(productID string) {
				defer wg.Done()
				productOrder := pgModel.ProductOrder{}
				for _, preProductOrder := range preProductOrders {
					if productID == (preProductOrder.ProductID).String() {
						productOrder = preProductOrder
						break
					}
				}
				errFunc := tx.Unscoped().Where("id = ?", productOrder.ID).Delete(&pgModel.ProductOrder{}).Error
				if errFunc != nil {
					err = errFunc
					return
				}

				// Update quantity product
				if order.Type == "export" {
					productRes, errFunc := s.productSvc.GetProduct(ctx, requestModel.GetProductRequest{
						ParamProductID: productID,
					})
					if errFunc != nil {
						err = errFunc
						return
					}
					product := productRes.Product
					updateQuantity := product.Quantity + productOrder.OrderQuantity
					productUpdateRes, errFunc := s.productSvc.UpdateProduct(ctx, requestModel.UpdateProductRequest{
						ParamProductID: (product.ID).String(),
						Quantity:       &updateQuantity,
					})
					if productUpdateRes != nil && productUpdateRes.TransactionID != nil {
						arrOtherServiceTX = append(arrOtherServiceTX, *productUpdateRes.TransactionID)
					}
					if errFunc != nil {
						err = errFunc
						return
					}
				}

			}(productID)
		}

		createProductIDs := helper.DifferenceArray(currentProductIDs, sameProductIDs)

		// Create product order
		for _, productID := range createProductIDs {
			wg.Add(1)
			go func(productID string) {
				defer wg.Done()
				productRes, errFunc := s.productSvc.GetProduct(ctx, requestModel.GetProductRequest{
					ParamProductID: productID,
				})
				if errFunc != nil {
					err = errFunc
					return
				}
				product := productRes.Product
				productIDUUID, _ := pgModel.UUIDFromString(productID)
				createOrderProduct := &pgModel.ProductOrder{}

				for _, productInfo := range req.OrderProductInfo {
					if productInfo.ProductID == productID {
						createOrderProduct = &pgModel.ProductOrder{
							ProductID:      &productIDUUID,
							OrderID:        &order.ID,
							OrderQuantity:  productInfo.OrderQuantity,
							OrderRealPrice: productInfo.OrderRealPrice,
						}
						break
					}
				}
				errFunc = tx.Create(createOrderProduct).Error
				if errFunc != nil {
					err = errFunc
					return
				}

				checkUpdateProduct := false
				if order.Type == "export" {
					if product.Quantity < createOrderProduct.OrderQuantity {
						err = errors.InvalidOrderProductQuantityError
						return
					}
					product.Quantity -= createOrderProduct.OrderQuantity
					checkUpdateProduct = true
				} else if order.Type == "import" && req.Status == "done" {
					product.Quantity += createOrderProduct.OrderQuantity
					checkUpdateProduct = true
				}

				// Update product quantity
				if checkUpdateProduct {
					updateQuantity := product.Quantity
					updateProductReq := requestModel.UpdateProductRequest{
						ParamProductID: (product.ID).String(),
						Quantity:       &updateQuantity,
					}
					updateProductRes, errFunc := s.productSvc.UpdateProduct(ctx, updateProductReq)

					if updateProductRes != nil && updateProductRes.TransactionID != nil {
						arrOtherServiceTX = append(arrOtherServiceTX, *updateProductRes.TransactionID)
					}
					if errFunc != nil {
						err = errFunc
						return
					}
					product = updateProductRes.Product
				}

				orderProductInfo = append(orderProductInfo, &responseModel.OrderProductInfo{
					Product: &pgModel.Product{
						Model: pgModel.Model{
							ID: product.ID,
						},
						Name:      product.Name,
						Price:     product.Price,
						MainPhoto: product.MainPhoto,
						Quantity:  product.Quantity,
					},
					OrderQuantity:  createOrderProduct.OrderQuantity,
					OrderRealPrice: createOrderProduct.OrderRealPrice,
				})
			}(productID)
		}

		arrSameProductOrder := []pgModel.ProductOrder{}

		type updateProductOrder struct {
			pre    pgModel.ProductOrder
			update *requestModel.ProductOrder
		}
		arrUpdateProductOrder := []updateProductOrder{}

		for _, item := range sameProductIDs {
			var pre pgModel.ProductOrder
			var update *requestModel.ProductOrder
			for _, value := range preProductOrders {
				if item == (value.ProductID).String() {
					pre = value
					break
				}
			}

			for _, value := range req.OrderProductInfo {
				if item == value.ProductID && (value.OrderQuantity != pre.OrderQuantity || value.OrderRealPrice != pre.OrderRealPrice) {
					update = &value
					break
				}
			}

			if update != nil {
				arrUpdateProductOrder = append(arrUpdateProductOrder, updateProductOrder{
					pre:    pre,
					update: update,
				})
			} else {
				arrSameProductOrder = append(arrSameProductOrder, pre)
			}
		}

		for _, item := range arrSameProductOrder {
			wg.Add(1)
			go func(productOrder pgModel.ProductOrder) {
				defer wg.Done()
				productRes, errFunc := s.productSvc.GetProduct(ctx, requestModel.GetProductRequest{
					ParamProductID: (productOrder.ProductID).String(),
				})
				if errFunc != nil {
					err = errFunc
					return
				}
				product := productRes.Product
				if order.Type == "import" && req.Status == "done" {
					quantityUpdate := product.Quantity + productOrder.OrderQuantity
					productUpdateReq := requestModel.UpdateProductRequest{
						ParamProductID: (productOrder.ProductID).String(),
						Quantity:       &quantityUpdate,
					}
					updateProductRes, errFunc := s.productSvc.UpdateProduct(ctx, productUpdateReq)
					if updateProductRes != nil && updateProductRes.TransactionID != nil {
						arrOtherServiceTX = append(arrOtherServiceTX, *updateProductRes.TransactionID)
					}
					if errFunc != nil {
						err = errFunc
						return
					}
					product = updateProductRes.Product
				}
				orderProductInfo = append(orderProductInfo, &responseModel.OrderProductInfo{
					Product: &pgModel.Product{
						Model: pgModel.Model{
							ID: product.ID,
						},
						Name:      product.Name,
						Price:     product.Price,
						MainPhoto: product.MainPhoto,
						Quantity:  product.Quantity,
					},
					OrderQuantity:  productOrder.OrderQuantity,
					OrderRealPrice: productOrder.OrderRealPrice,
				})
			}(item)
		}

		for _, item := range arrUpdateProductOrder {
			wg.Add(1)
			go func(updateProductOrder updateProductOrder) {
				defer wg.Done()
				updateRecord := updateProductOrder.pre
				updateRecord.OrderQuantity = updateProductOrder.update.OrderQuantity
				updateRecord.OrderRealPrice = updateProductOrder.update.OrderRealPrice
				errFunc := tx.Save(updateRecord).Error
				if errFunc != nil {
					err = errFunc
					return
				}

				productRes, errFunc := s.productSvc.GetProduct(ctx, requestModel.GetProductRequest{
					ParamProductID: (updateRecord.ProductID).String(),
				})
				if errFunc != nil {
					err = errFunc
					return
				}
				product := productRes.Product

				if (req.Status == "done" && order.Type == "import") || order.Type == "export" {

					productUpdateReq := requestModel.UpdateProductRequest{
						ParamProductID: (product.ID).String(),
					}
					if req.Status == "done" && order.Type == "import" {
						quantityUpdate := product.Quantity + updateRecord.OrderQuantity
						productUpdateReq.Quantity = &quantityUpdate

					} else {
						durationQuantityUpdate := updateProductOrder.update.OrderQuantity - updateProductOrder.pre.OrderQuantity

						if product.Quantity < durationQuantityUpdate {
							err = errors.InvalidOrderProductQuantityError
							return
						}
						updateRecord := updateProductOrder.pre
						updateRecord.OrderQuantity = updateProductOrder.update.OrderQuantity
						updateRecord.OrderRealPrice = updateProductOrder.update.OrderRealPrice

						errFunc = tx.Save(updateRecord).Error
						if errFunc != nil {
							err = errFunc
							return
						}
						quantityUpdate := product.Quantity - durationQuantityUpdate
						productUpdateReq.Quantity = &quantityUpdate
					}
					updateProductRes, errFunc := s.productSvc.UpdateProduct(ctx, productUpdateReq)
					if updateProductRes != nil && updateProductRes.TransactionID != nil {
						arrOtherServiceTX = append(arrOtherServiceTX, *updateProductRes.TransactionID)
					}
					if errFunc != nil {
						err = errFunc
						return
					}
					product = updateProductRes.Product
				}
				orderProductInfo = append(orderProductInfo, &responseModel.OrderProductInfo{
					Product: &pgModel.Product{
						Model: pgModel.Model{
							ID: product.ID,
						},
						Name:      product.Name,
						Price:     product.Price,
						MainPhoto: product.MainPhoto,
						Quantity:  product.Quantity,
					},
					OrderQuantity:  updateRecord.OrderQuantity,
					OrderRealPrice: updateRecord.OrderRealPrice,
				})
			}(item)
		}

		wg.Wait()

		if err != nil {
			s.rollbackTransactions(ctx, arrOtherServiceTX)
			return res, err
		}
	}

	if req.ReceiverPhoneNumber != "" && req.ReceiverPhoneNumber != order.ReceiverPhoneNumber {
		order.ReceiverPhoneNumber = req.ReceiverPhoneNumber
	}

	if req.ReceiverAddress != "" && req.ReceiverAddress != order.ReceiverAddress {
		order.ReceiverAddress = req.ReceiverAddress
	}

	if req.ReceiverFullname != "" && req.ReceiverFullname != order.ReceiverFullname {
		order.ReceiverFullname = req.ReceiverFullname
	}
Update:
	err = tx.Save(order).Error

	if err != nil {
		s.rollbackTransactions(ctx, arrOtherServiceTX)
		return res, err
	}

	s.commitTransactions(ctx, arrOtherServiceTX)

	orderRes, err := s.getOrderDetail(ctx, *order, false)

	if err != nil {
		return res, err
	}

	res.Order = orderRes

	if order.ImplementerID != nil {
		implementer := &pgModel.User{
			Model: pgModel.Model{
				ID: *order.ImplementerID,
			},
		}
		s.db.Find(implementer, implementer)
		res.Order.Implementer = &pgModel.User{
			Model: pgModel.Model{
				ID: implementer.ID,
			},
			Fullname: implementer.Fullname,
			Role:     implementer.Role,
		}
	}

	return res, nil
}

func (s *pgService) getOrderDetail(ctx context.Context, order pgModel.Order, isGetProductInfo bool) (*responseModel.Order, error) {
	orderRes := &responseModel.Order{}

	var err error

	var creator *pgModel.User
	getCreatorReq := requestModel.GetUserRequest{
		ParamUserID: (order.CreatorID).String(),
	}

	var (
		implementer *pgModel.User
		customer    *pgModel.User
	)

	orderProductInfo := []*responseModel.OrderProductInfo{}

	wg := sync.WaitGroup{}

	// Get customer
	if order.Type != "import" {
		getCustomerReq := requestModel.GetUserRequest{
			ParamUserID: (order.CustomerID).String(),
		}

		wg.Add(1)
		go func(getCustomerReq requestModel.GetUserRequest) {
			defer wg.Done()
			customerRes, errFunc := s.userSvc.GetUser(ctx, getCustomerReq)
			if errFunc != nil {
				err = errFunc
				return
			}
			customer = customerRes.User
		}(getCustomerReq)
	}

	// Get creator
	wg.Add(1)
	go func(getCreatorReq requestModel.GetUserRequest) {
		defer wg.Done()
		creatorRes, errFunc := s.userSvc.GetUser(ctx, getCreatorReq)
		if errFunc != nil {
			err = errFunc
			return
		}
		creator = &pgModel.User{
			Model: pgModel.Model{
				ID: creatorRes.User.ID,
			},
			Fullname: creatorRes.User.Fullname,
			Role:     creatorRes.User.Role,
		}
	}(getCreatorReq)

	// Get implementer
	if order.Status == "done" {
		getImplementerReq := requestModel.GetUserRequest{
			ParamUserID: (order.ImplementerID).String(),
		}
		wg.Add(1)
		go func(getImplementerReq requestModel.GetUserRequest) {
			defer wg.Done()
			implementerRes, errFunc := s.userSvc.GetUser(ctx, getImplementerReq)
			if errFunc != nil {
				err = errFunc
				return
			}
			implementer = &pgModel.User{
				Model: pgModel.Model{
					ID: implementerRes.User.ID,
				},
				Fullname: implementerRes.User.Fullname,
				Role:     implementerRes.User.Role,
			}
		}(getImplementerReq)
	}

	// Get product order
	if isGetProductInfo {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// implementer := &pgModel.U

			productOrders := []pgModel.ProductOrder{}
			errFunc := s.db.Where("order_id = ?", (order.ID).String()).Find(&productOrders).Error
			if errFunc != nil {
				err = errFunc
				return
			}
			for _, productOrder := range productOrders {
				wg.Add(1)
				go func(productOrderInfo pgModel.ProductOrder) {
					defer wg.Done()
					getProductReq := requestModel.GetProductRequest{
						ParamProductID: (productOrderInfo.ProductID).String(),
					}
					productRes, errFunc := s.productSvc.GetProduct(ctx, getProductReq)
					if errFunc != nil {
						err = errFunc
						return
					}
					product := productRes.Product
					orderProductInfo = append(orderProductInfo, &responseModel.OrderProductInfo{
						OrderQuantity:  productOrderInfo.OrderQuantity,
						OrderRealPrice: productOrderInfo.OrderRealPrice,
						Product: &pgModel.Product{
							Model: pgModel.Model{
								ID: product.ID,
							},
							Name:      product.Name,
							Price:     product.Price,
							MainPhoto: product.MainPhoto,
							Quantity:  product.Quantity,
							Color:     product.Color,
						},
					})
				}(productOrder)
			}
		}()
	}

	wg.Wait()

	if err != nil {
		return orderRes, err
	}

	orderRes.Order = s.removeUnnecessaryOrderInfo(ctx, order)
	orderRes.Creator = creator
	orderRes.Customer = customer
	orderRes.Implementer = implementer
	orderRes.OrderProductInfo = orderProductInfo

	return orderRes, nil
}

func (s *pgService) getOrderCustomer(ctx context.Context, order pgModel.Order) (*pgModel.User, error) {
	res := &pgModel.User{}

	getCustomerReq := requestModel.GetUserRequest{
		ParamUserID: (order.CustomerID).String(),
	}

	customerRes, errFunc := s.userSvc.GetUser(ctx, getCustomerReq)
	if errFunc != nil {
		return res, nil
	}
	res = customerRes.User
	return res, nil
}

func (s *pgService) getOrderCreator(ctx context.Context, order pgModel.Order) (*pgModel.User, error) {
	res := &pgModel.User{}

	getCreatorReq := requestModel.GetUserRequest{
		ParamUserID: (order.CreatorID).String(),
	}

	creatorRes, errFunc := s.userSvc.GetUser(ctx, getCreatorReq)
	if errFunc != nil {
		return res, errFunc
	}
	res = &pgModel.User{
		Model: pgModel.Model{
			ID: creatorRes.User.ID,
		},
		Fullname: creatorRes.User.Fullname,
		Role:     creatorRes.User.Role,
	}

	return res, nil
}

func (s *pgService) getOrderImplementer(ctx context.Context, order pgModel.Order) (*pgModel.User, error) {
	res := &pgModel.User{}
	getImplementerReq := requestModel.GetUserRequest{
		ParamUserID: (order.ImplementerID).String(),
	}
	implementerRes, errFunc := s.userSvc.GetUser(ctx, getImplementerReq)
	if errFunc != nil {
		return res, errFunc
	}
	res = &pgModel.User{
		Model: pgModel.Model{
			ID: implementerRes.User.ID,
		},
		Fullname: implementerRes.User.Fullname,
		Role:     implementerRes.User.Role,
	}
	return res, nil
}

func (s *pgService) removeUnnecessaryOrderInfo(ctx context.Context, order pgModel.Order) *pgModel.Order {
	order.CreatorID = nil
	order.ImplementerID = nil
	order.CustomerID = nil
	return &order
}
