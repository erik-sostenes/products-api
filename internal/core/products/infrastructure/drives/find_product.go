package drives

import (
	"net/http"
	"strings"

	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
	"github.com/erik-sostenes/products-api/internal/shared/infrastructure/drives/handler"
	"github.com/labstack/echo/v4"
)

// FindProduct represents an http handler, it is in charge of getting the data by identifier to the core business
func FindProduct(query query.Bus[services.FindProductQuery, services.ProductResponse]) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("id")

		if strings.TrimSpace(id) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "Missing identifier."})
		}

		productQuery := services.FindProductQuery{
			Id: id,
		}

		product, err := query.Ask(c.Request().Context(), productQuery)

		if err != nil {
			return handler.Error(err)
		}

		return c.JSON(http.StatusOK, echo.Map{"data": product})
	}
}
