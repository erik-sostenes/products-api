package middlewareJWT

import (
	"fmt"
	"net/http"

	"github.com/erik-sostenes/products-api/internal/shared/infrastructure/drives/jwt"
	"github.com/labstack/echo/v4"
)

// Authorization middleware that validates the token for each http request
// if the token is invalid the client is responded to with a StatusForbidden
//
// if the token is valid the requested HandlerFunc is executed
func Authorization(jwt jwt.Token[jwt.Claims], next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")

		err := jwt.Validate(token)
		if err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusForbidden, echo.Map{"error": "Prohibited from accessing this resource."})
		}

		return next(c)
	}
}
