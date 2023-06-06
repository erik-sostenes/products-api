package services

import (
	"context"

	"github.com/erik-sostenes/products-api/internal/core/products/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/products/business/ports"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/command"
)

// ProductCommand implements command.Command interface
var _ command.Command = ProductCommand{}

type SalesCommand struct {
	Days   int
	Amount float64
}

// ProductCommand represents a DTO(Data Transfer Object)
type ProductCommand struct {
	Id        string
	Title     string
	ImageURL  string
	Price     float64
	Rating    float64
	Offer     bool
	Available bool
	Sales     SalesCommand
}

// CommandId returns an identifier of ProductCommand
func (ProductCommand) CommandId() string {
	return "create_product_command"
}

// CreateProductCommandHandler implements command.Handler interface
var _ command.Handler[ProductCommand] = &CreateProductCommandHandler{}

type CreateProductCommandHandler struct {
	ports.ProductStorer
}

// NewCreateProductCommandHandler returns an instance of CreateProductCommandHandler
func NewCreateProductCommandHandler(store ports.ProductStorer) CreateProductCommandHandler {
	return CreateProductCommandHandler{
		ProductStorer: store,
	}
}

// Handler method that receives the command, the command is converted into an Entity and sent to the adapter that implements the ports.ProductStorer interface
func (c *CreateProductCommandHandler) Handler(ctx context.Context, cmd ProductCommand) error {
	product, err := domain.NewProduct(cmd.Id, cmd.Title, cmd.ImageURL, cmd.Price, cmd.Rating, cmd.Offer, cmd.Available, cmd.Sales.Days, cmd.Sales.Amount)
	if err != nil {
		return err
	}

	return c.ProductStorer.Save(ctx, product.Id, product)
}
