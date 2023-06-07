package drives

import (
	"net/http"
	"strings"

	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/shared/domain"
	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
	"github.com/labstack/echo/v4"
)

// FindProduct represents an http handler, it is in charge of getting the data by identifier to the core business
func FindProduct(services services.FinderProduct) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("id")

		if strings.TrimSpace(id) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, domain.Map{"error": "Missing identifier."})
		}

		product, err := services.FindById(c.Request().Context(), id)

		switch err.(type) {
		case wrongs.StatusBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		case wrongs.StatusNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			echo.NewHTTPError(http.StatusInternalServerError, domain.Map{"error": "An error occurred in the http server."})
		}

		return c.JSON(http.StatusOK, domain.Map{"data": product})
	}
}
