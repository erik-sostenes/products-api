package drives

import (
	"net/http"
	"strings"

	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/command"
	"github.com/erik-sostenes/products-api/internal/shared/infrastructure/drives/handler"
	"github.com/labstack/echo/v4"
)

// DeleteProduct represents an http handler, it is in charge of validating the http response
// and sending the data to the core business to delete a record
func DeleteProduct(cmdBus command.CommandBus[services.DeleteProductCommand]) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		if strings.TrimSpace(id) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "Missing identifier."})
		}

		cmd := services.DeleteProductCommand{
			Id: id,
		}

		if err := cmdBus.Dispatch(c.Request().Context(), cmd); err != nil {
			return handler.Error(err)
		}

		return c.JSON(http.StatusOK, echo.Map{"message": "ok"})
	}
}
