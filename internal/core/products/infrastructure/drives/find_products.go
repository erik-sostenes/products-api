package drives

import (
	"net/http"

	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/shared/domain"
	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
	"github.com/labstack/echo/v4"
)

// FindProduct represents an http handler, it is in charge of getting the data to the core business
func FindProducts(services services.FinderProducts) echo.HandlerFunc {
	return func(c echo.Context) error {
		products, err := services.ProductStorer.Find(c.Request().Context())

		switch err.(type) {
		case wrongs.StatusBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			echo.NewHTTPError(http.StatusInternalServerError, domain.Map{"error": "An error occurred in the http server."})
		}

		return c.JSON(http.StatusOK, domain.Map{"data": products})
	}
}
