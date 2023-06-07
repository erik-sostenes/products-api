// bootstrap package that builds the program with its full set of components
package bootstrap

import (
	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/core/products/infrastructure/driven/db"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/command"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
	"github.com/labstack/echo/v4"
)

// NewInjector injects all the dependencies to start the app
func NewInjector() error {
	engine := echo.New()

	mock := db.NewMockProductStorer()

	productCommandHandler := services.NewCreateProductCommandHandler(mock)
	createProductCmdBus := make(command.CommandBus[services.ProductCommand])
	if err := createProductCmdBus.Record(services.ProductCommand{}, &productCommandHandler); err != nil {
		return err
	}

	deleteProductCommandHandler := services.NewDeleteProductCommandHandler(mock)
	deleteProductCmdBus := make(command.CommandBus[services.DeleteProductCommand])
	if err := deleteProductCmdBus.Record(services.DeleteProductCommand{}, deleteProductCommandHandler); err != nil {
		return err
	}

	finderProductsQueryHandler := services.NewFinderProductsQueryHandler(mock)
	productsQueryBus := make(query.QueryBus[services.FindProductsQuery, []services.ProductResponse])
	if err := productsQueryBus.Record(services.FindProductsQuery{}, finderProductsQueryHandler); err != nil {
		return err
	}

	finderProductQueryHandler := services.NewFinderProductQueryHandler(mock)
	productQueryBus := make(query.QueryBus[services.FindProductQuery, services.ProductResponse])
	if err := productQueryBus.Record(services.FindProductQuery{}, finderProductQueryHandler); err != nil {
		return err
	}

	return NewServer(engine, createProductCmdBus, deleteProductCmdBus, productsQueryBus, productQueryBus).Start()
}
