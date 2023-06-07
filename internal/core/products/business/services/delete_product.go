package services

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/products/business/ports"
	"github.com/erik-sostenes/products-api/internal/shared/domain"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/command"
)

// DeleteProductCommand implements command.Command interface
var _ command.Command = DeleteProductCommand{}

// DeleteProductCommand represents an DTO and the action to be performed
type DeleteProductCommand struct {
	Id string
}

// CommandId returns an identifier of ProductCommand
func (DeleteProductCommand) CommandId() string {
	return "delete_product_command"
}

// DeleteProductCommandHandler implements command.Handler interface
var _ command.Handler[DeleteProductCommand] = &DeleteProductCommandHandler{}

type DeleteProductCommandHandler struct {
	ports.ProductStorer
}

// NewDeleteProductCommandHandler returns an instance of command.Handler with injected dependencies
func NewDeleteProductCommandHandler(store ports.ProductStorer) command.Handler[DeleteProductCommand] {
	return &DeleteProductCommandHandler{
		ProductStorer: store,
	}
}

// Handler method that receives the command, the command is converted into an Object Domain and sent to the adapter that implements the ports.ProductStorer interface
func (d *DeleteProductCommandHandler) Handler(ctx context.Context, cmd DeleteProductCommand) error {
	if err := domain.UuID(cmd.Id).EnsureIdIsValid(); err != nil {
		return err
	}

	return d.ProductStorer.Delete(ctx, cmd.Id)
}
