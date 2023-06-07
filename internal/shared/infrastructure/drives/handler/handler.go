package handler

import (
	"net/http"

	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
	"github.com/labstack/echo/v4"
)

// HandlerError handles the type of error
//
// responds with http status code  along with the message
func Error(err error) error {
	switch err.(type) {
	case wrongs.StatusBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err})
	case wrongs.StatusNotFound:
		return echo.NewHTTPError(http.StatusNotFound, echo.Map{"error": err})
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"error": err})
	}
}
