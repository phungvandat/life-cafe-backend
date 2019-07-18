package product

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	requesetModel "github.com/phungvandat/life-cafe-backend/model/request"
	"github.com/phungvandat/life-cafe-backend/service"
)

// MakeCreateProductEndpoint make endpoint for create product
func MakeCreateProductEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.CreateProductRequest)
		res, err := s.ProductService.CreateProduct(ctx, req)
		var errTransaction error
		var transactionID *string
		if res != nil && res.TransactionID != nil {
			transactionID = res.TransactionID
		}
		if err != nil {
			if transactionID != nil {
				errTransaction = s.ProductService.RollbackTransaction(ctx, *transactionID)
			}
			return nil, err
		}
		if transactionID != nil {
			errTransaction = s.ProductService.CommitTransaction(ctx, *transactionID)
			res.TransactionID = nil
		}
		if errTransaction != nil {
			var log log.Logger
			log.Log("Transaction create product failure by error ", errTransaction)
		}
		return res, nil
	}
}

// MakeGetProductEndpoint func
func MakeGetProductEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.GetProductRequest)
		res, err := s.ProductService.GetProduct(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// MakeGetProductsEndpoint func
func MakeGetProductsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.GetProductsRequest)

		res, err := s.ProductService.GetProducts(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// MakeUpdateProductEndpoint func
func MakeUpdateProductEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requesetModel.UpdateProductRequest)

		res, err := s.ProductService.UpdateProduct(ctx, req)

		var errTransaction error
		var transactionID *string
		if res != nil && res.TransactionID != nil {
			transactionID = res.TransactionID
		}
		if err != nil {
			if transactionID != nil {
				errTransaction = s.ProductService.RollbackTransaction(ctx, *transactionID)
			}
			return nil, err
		}
		if transactionID != nil {
			errTransaction = s.ProductService.CommitTransaction(ctx, *transactionID)
			res.TransactionID = nil
		}
		if errTransaction != nil {
			var log log.Logger
			log.Log("Transaction create product failure by error ", errTransaction)
		}
		return res, nil
	}
}
