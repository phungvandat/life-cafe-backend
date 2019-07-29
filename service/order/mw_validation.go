package order

import (
	"context"
	"strconv"
	"strings"

	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
	requestModel "github.com/phungvandat/life-cafe-backend/model/request"
	responseModel "github.com/phungvandat/life-cafe-backend/model/response"
	"github.com/phungvandat/life-cafe-backend/util/constants"
	errors "github.com/phungvandat/life-cafe-backend/util/error"
	"github.com/phungvandat/life-cafe-backend/util/regex"
)

type validationMiddleware struct {
	Service
}

// ValidationMiddleware for validation purposes
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) RollbackTransaction(ctx context.Context, transactionID string) error {
	return mw.Service.RollbackTransaction(ctx, transactionID)
}

func (mw validationMiddleware) CommitTransaction(ctx context.Context, transactionID string) error {
	return mw.Service.CommitTransaction(ctx, transactionID)
}

func (mw validationMiddleware) CreateOrder(ctx context.Context, req requestModel.CreateOrderRequest) (*responseModel.CreateOrderResponse, error) {
	if req.Type == "" {
		return nil, errors.OrderTypeIsRequiredError
	}

	if _, check := constants.OrderType[req.Type]; !check {
		return nil, errors.InvalidOrderTypeError
	}

	if req.Status == "" {
		return nil, errors.OrderStatusIsRequiredError
	}

	if _, check := constants.OrderStatus[req.Status]; !check {
		return nil, errors.InvalidOrderStatusError
	}

	if req.OrderProductInfo == nil || len(req.OrderProductInfo) <= 0 {
		return nil, errors.OrderProductInfoIsRequiredError
	}

	for _, productInfo := range req.OrderProductInfo {
		if productInfo.ProductID == "" {
			return nil, errors.ProductIDIsRequiredError
		}
		if _, err := pgModel.UUIDFromString(productInfo.ProductID); err != nil {
			return nil, errors.InvalidProductIDTypeError
		}

		if productInfo.OrderQuantity <= 0 {
			return nil, errors.InvalidOrderProductQuantityError
		}

		if productInfo.OrderRealPrice < 0 {
			return nil, errors.InvalidOrderProductOrderRealPriceError
		}
	}

	if req.Type != "import" {
		if req.ReceiverPhoneNumber == "" {
			return nil, errors.ReceiverPhoneNumberIsRequiredError
		}

		if !regex.IsPhoneNumberValid(req.ReceiverPhoneNumber) {
			return nil, errors.InvalidReceiverPhoneNumberError
		}

		if strings.Trim(req.ReceiverAddress, " ") == "" {
			return nil, errors.ReceiverAddressIsRequiredError
		}

		if strings.Trim(req.ReceiverFullname, " ") == "" {
			return nil, errors.ReceiverFullnameIsRequiredError
		}

		if req.CustomerID != "" {
			if _, err := pgModel.UUIDFromString(req.CustomerID); err != nil {
				return nil, errors.InvalidCustomerIDError
			}
		} else {
			if req.CustomerPhoneNumber == "" {
				return nil, errors.CustomerPhoneNumberIsReqiredError
			}
			if !regex.IsPhoneNumberValid(req.CustomerPhoneNumber) {
				return nil, errors.InvalidCustomerPhoneNumberError
			}
			if strings.Trim(req.CustomerAddress, " ") == "" {
				return nil, errors.CustomerAddressIsRequiredError
			}
			if strings.Trim(req.CustomerFullname, " ") == "" {
				return nil, errors.CustomerFullnameIsRequiredError
			}
		}
	}

	return mw.Service.CreateOrder(ctx, req)
}

func (mw validationMiddleware) GetOrders(ctx context.Context, req requestModel.GetOrdersRequest) (*responseModel.GetOrdersResponse, error) {

	if _, err := strconv.ParseInt(req.Skip, 10, 32); req.Skip != "" && err != nil {
		return nil, errors.InvalidSkipError
	}

	if _, err := strconv.ParseInt(req.Limit, 10, 32); req.Limit != "" && err != nil {
		return nil, errors.InvalidLimitError
	}

	return mw.Service.GetOrders(ctx, req)
}

func (mw validationMiddleware) UpdateOrder(ctx context.Context, req requestModel.UpdateOrderRequest) (*responseModel.UpdateOrderResponse, error) {

	if _, check := constants.OrderStatus[req.Status]; req.Status != "" && !check {
		return nil, errors.InvalidOrderStatusError
	}

	if req.ReceiverPhoneNumber == "" && !regex.IsPhoneNumberValid(req.ReceiverPhoneNumber) {
		return nil, errors.InvalidReceiverPhoneNumberError
	}

	if len(req.OrderProductInfo) > 0 {
		for _, productInfo := range req.OrderProductInfo {
			if productInfo.ProductID == "" {
				return nil, errors.ProductIDIsRequiredError
			}
			if _, err := pgModel.UUIDFromString(productInfo.ProductID); err != nil {
				return nil, errors.InvalidProductIDTypeError
			}

			if productInfo.OrderQuantity <= 0 {
				return nil, errors.InvalidOrderProductQuantityError
			}

			if productInfo.OrderRealPrice < 0 {
				return nil, errors.InvalidOrderProductOrderRealPriceError
			}
		}
	}

	return mw.Service.UpdateOrder(ctx, req)
}
