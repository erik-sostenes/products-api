package drives

import (
	"net/http"

	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
	"github.com/erik-sostenes/products-api/internal/shared/infrastructure/drives/handler"
	"github.com/labstack/echo/v4"
)

// FindProduct represents an http handler, it is in charge of getting the data to the core business
func FindProducts(query query.Bus[services.FindProductsQuery, []services.ProductResponse]) echo.HandlerFunc {
	return func(c echo.Context) error {
		productsQuery := services.FindProductsQuery{}

		products, err := query.Ask(c.Request().Context(), productsQuery)

		if len(products) == 0 {
			return c.JSON(http.StatusNoContent, echo.Map{"message": "No content of products."})
		}

		if err != nil {
			return handler.Error(err)
		}

		return c.JSON(http.StatusOK, echo.Map{"data": products})
	}
}
