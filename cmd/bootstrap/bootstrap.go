package bootstrap

import (
	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/core/products/infrastructure/driven/db"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/command"
	"github.com/labstack/echo/v4"
)

// NewInjector injects all the dependencies to start the app
func NewInjector() error {
	engine := echo.New()

	productCommandHandler := services.NewCreateProductCommandHandler(db.NewMockProductStorer())

	cmdBus := make(command.CommandBus[services.ProductCommand])
	if err := cmdBus.Record(services.ProductCommand{}, &productCommandHandler); err != nil {
		return err
	}

	return NewServer(engine, cmdBus).Start()
}
