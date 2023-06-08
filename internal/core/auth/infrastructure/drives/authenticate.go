package drives

import (
	"net/http"
	"strings"

	"github.com/erik-sostenes/products-api/internal/core/auth/business/domain/wrongs"
	"github.com/erik-sostenes/products-api/internal/core/auth/business/services"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
	"github.com/erik-sostenes/products-api/internal/shared/infrastructure/drives/handler"
	"github.com/labstack/echo/v4"
)

func Authenticate(bus query.Bus[services.AuthenticateAccountQuery, services.AuthResponse]) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("id")
		if strings.TrimSpace(id) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "Missing query parameter id"})
		}

		password := c.QueryParam("password")
		if strings.TrimSpace(id) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "Missing query parameter password"})
		}

		query := services.AuthenticateAccountQuery{
			Id:       id,
			Password: password,
		}

		authResponse, err := bus.Ask(c.Request().Context(), query)
		if err != nil {
			switch err.(type) {
			case wrongs.InvalidAuthAccount:
				return echo.NewHTTPError(http.StatusForbidden, echo.Map{"error": err.Error()})
			case wrongs.InvalidAuthCredentials:
				return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{"error": err.Error()})
			default:
				return handler.Error(err)
			}
		}

		return c.JSON(http.StatusOK, echo.Map{"response": authResponse})
	}
}
