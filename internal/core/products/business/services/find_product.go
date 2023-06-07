package services

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/products/business/ports"
	"github.com/erik-sostenes/products-api/internal/shared/domain"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
)

var _ query.Query = FindProductQuery{}

type FindProductQuery struct {
	Id string
}

func (FindProductQuery) QueryId() string {
	return "find_product_query"
}

var _ query.Handler[FindProductQuery, ProductResponse] = FinderProductQueryHandler{}

type FinderProductQueryHandler struct {
	ports.ProductStorer
}

func NewFinderProductQueryHandler(store ports.ProductStorer) query.Handler[FindProductQuery, ProductResponse] {
	return FinderProductQueryHandler{
		ProductStorer: store,
	}
}

// Handler method that receives the request to obtain a product record by identifier, requesting it from the adapter that implements the ports.ProductStorer interface
func (f FinderProductQueryHandler) Handler(ctx context.Context, query FindProductQuery) (ProductResponse, error) {
	if err := domain.UuID(query.Id).EnsureIdIsValid(); err != nil {
		return ProductResponse{}, err
	}

	product, err := f.ProductStorer.FindById(ctx, query.Id)

	if err != nil {
		return ProductResponse{}, err
	}

	return ProductResponse(product), err

}
