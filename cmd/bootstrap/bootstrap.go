package bootstrap

import "github.com/labstack/echo/v4"

// NewInjector injects all the dependencies to start the app
func NewInjector() (err error) {
	engine := echo.New()

	return NewServer(engine).Start()
}
