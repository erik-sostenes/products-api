package services

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/products/business/ports"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
)

var _ query.Query = FindProductsQuery{}

type FindProductsQuery struct {
	Id string
}

func (FindProductsQuery) QueryId() string {
	return "find_products_query"
}

var _ query.Handler[FindProductsQuery, []ProductResponse] = FinderProductsQueryHandler{}

type FinderProductsQueryHandler struct {
	ports.ProductStorer
}

func NewFinderProductsQueryHandler(store ports.ProductStorer) query.Handler[FindProductsQuery, []ProductResponse] {
	return FinderProductsQueryHandler{
		ProductStorer: store,
	}
}

// Handler method that receives the request to get all product records, requesting it to the adapter implementing the ports.ProductStorer interface
func (f FinderProductsQueryHandler) Handler(ctx context.Context, _ FindProductsQuery) ([]ProductResponse, error) {
	products, err := f.ProductStorer.Find(ctx)
	if err != nil {
		return []ProductResponse{}, err
	}

	var productsResponse []ProductResponse

	for _, product := range products {
		productsResponse = append(productsResponse, ProductResponse(product))
	}

	return productsResponse, err
}
