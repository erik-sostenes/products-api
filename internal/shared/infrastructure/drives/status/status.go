package status

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck http handler that checks the status of the server
func HealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}
}
