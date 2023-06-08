package bootstrap

import (
	"fmt"
	"os"
	"strings"

	servicesAuth "github.com/erik-sostenes/products-api/internal/core/auth/business/services"
	drivesAuth "github.com/erik-sostenes/products-api/internal/core/auth/infrastructure/drives"
	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/core/products/infrastructure/drives"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/command"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
	"github.com/erik-sostenes/products-api/internal/shared/infrastructure/drives/status"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = "8080"

// Server contains the configuration of server to start and register all http handler
type Server struct {
	port                   string
	engine                 *echo.Echo
	createProductCmdBus    command.CommandBus[services.ProductCommand]
	deleteProductCmdBus    command.CommandBus[services.DeleteProductCommand]
	finderProductsQueryBus query.QueryBus[services.FindProductsQuery, []services.ProductResponse]
	finderProductQueryBus  query.QueryBus[services.FindProductQuery, services.ProductResponse]
	authQueryBus           query.QueryBus[servicesAuth.AuthenticateAccountQuery, servicesAuth.AuthResponse]
}

// NewServer returns an instance of Server
func NewServer(
	engine *echo.Echo,
	createProductCmdBus command.CommandBus[services.ProductCommand],
	deleteProductCmdBus command.CommandBus[services.DeleteProductCommand],
	finderProductsQueryBus query.QueryBus[services.FindProductsQuery, []services.ProductResponse],
	finderProductQueryBus query.QueryBus[services.FindProductQuery, services.ProductResponse],
	authQueryBus query.QueryBus[servicesAuth.AuthenticateAccountQuery, servicesAuth.AuthResponse],
) *Server {
	port := os.Getenv("PORT")
	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}

	return &Server{
		port:                   port,
		engine:                 engine,
		createProductCmdBus:    createProductCmdBus,
		deleteProductCmdBus:    deleteProductCmdBus,
		finderProductsQueryBus: finderProductsQueryBus,
		finderProductQueryBus:  finderProductQueryBus,
		authQueryBus:           authQueryBus,
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

	s.engine.GET("/api/v1/authenticate/", drivesAuth.Authenticate(&s.authQueryBus))

	group := s.engine.Group("/api/v1/products")

	group.GET("/status", status.HealthCheck())
	group.PUT("/create/:id", drives.CreateProduct(s.createProductCmdBus))
	group.GET("/get-all", drives.FindProducts(&s.finderProductsQueryBus))
	group.GET("/get/", drives.FindProduct(&s.finderProductQueryBus))
	group.DELETE("/delete/:id", drives.DeleteProduct(s.deleteProductCmdBus))
}
