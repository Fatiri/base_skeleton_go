package src

import (
	"fmt"
	"log"

	"github.com/base_skeleton_go/shared/tracer"
	v1 "github.com/base_skeleton_go/src/delivery/v1"
	"github.com/go-playground/validator"

	"github.com/base_skeleton_go/config"
	"github.com/base_skeleton_go/registry"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echoMiddleware "github.com/labstack/echo/middleware"
)

// Server ...
type Server struct {
	httpServer *echo.Echo
	uc         registry.UsecaseRegistry
	config     *config.Config
}

// InitServer ...
func InitServer(cfg *config.Config) *Server {
	echoServer := echo.New()
	echoServer.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}",` +
			`"latency_human":"${latency_human}\n"`,
	}))
	echoServer.Use(middleware.RequestID())

	// set tracer
	tracer.InitOpenTracing()

	echoServer.Validator = &CustomValidator{validator: validator.New()}
	registry := registry.NewUsecaseRegistry(*cfg)

	return &Server{httpServer: echoServer, uc: registry, config: cfg}
}

// Run ...
func (s *Server) Run() {
	accountDelivery := v1.NewAccountDelivery(s.uc.Account())
	accountGroup := s.httpServer.Group(s.config.ServiceName()+`/v1/account`, tracer.EchoRestTracerMiddleware)
	accountDelivery.Mount(accountGroup, s.config.Token().PublicKey)


	if err := s.httpServer.Start(fmt.Sprintf(":%d", s.config.Port())); err != nil {
		log.Panic(err)
	}
}

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
