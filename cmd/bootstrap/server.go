package bootstrap

import (
	"fmt"
	"os"
	"strings"

	"github.com/erik-sostenes/products-api/internal/shared/infrastructure/drives/status"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = "8080"

// Server contains the configuration of server to start and register all http handler
type Server struct {
	port   string
	engine *echo.Echo
}

// NewServer returns an instance of Server
func NewServer(engine *echo.Echo) *Server {
	port := os.Getenv("PORT")
	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}

	return &Server{
		port:   port,
		engine: engine,
	}
}

// Start initialize the server with all http handler
func (s *Server) Start() error {
	s.setRoutes()

	return s.engine.Start(fmt.Sprintf(":%v", s.port))
}

// Routes register all endpoints
//
// configure the middlewares CORS, Logger and Recover
func (s *Server) setRoutes() {
	s.engine.Use(middleware.CORS(), middleware.Recover(), middleware.Logger())
	group := s.engine.Group("/api/v1/products")

	group.GET("/status", status.HealthCheck())
}
