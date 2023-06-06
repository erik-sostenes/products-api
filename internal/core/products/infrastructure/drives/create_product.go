package drives

import (
	"fmt"
	"net/http"

	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/shared/domain"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/command"
	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
	"github.com/labstack/echo/v4"
)

type Sales struct {
	Days   int     `json:"days"`
	Amount float64 `json:"amount"`
}

// ProductRequest represents an http request
type ProductRequest struct {
	Id        string  `param:"id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	ImageURL  string  `json:"image_url"`
	Offer     bool    `json:"offer"`
	Available bool    `json:"available"`
	Rating    float64 `json:"rating"`
	Sales     Sales   `json:"sales"`
}

// CreateProduct represents an http handler, it is in charge of validating the http response and sending the data to the core business
func CreateProduct(bus command.CommandBus[services.ProductCommand]) echo.HandlerFunc {
	fmt.Println("Hola")
	return func(c echo.Context) error {
		var request ProductRequest

		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, domain.Map{"error": "The product structure is incorrect."})
		}

		cmd := services.ProductCommand{
			Id:        request.Id,
			Title:     request.Title,
			Price:     request.Price,
			ImageURL:  request.ImageURL,
			Offer:     request.Offer,
			Available: request.Available,
			Rating:    request.Rating,
			Sales: services.SalesCommand{
				Days:   request.Sales.Days,
				Amount: request.Sales.Amount,
			},
		}

		err := bus.Dispatch(c.Request().Context(), cmd)
		switch err.(type) {
		case wrongs.StatusBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, domain.Map{"message": err.Error()})
		default:
			echo.NewHTTPError(http.StatusInternalServerError, domain.Map{"error": "An error occurred in the http server."})
		}

		return c.JSON(http.StatusCreated, domain.Map{"message": "Product created."})
	}
}
