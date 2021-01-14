package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type IServer interface {
	Start() error
}

type Server struct {
	Port string
	Echo *echo.Echo
}

func NewServer(config Config) (*Server, error) {
	// Create Echo Server
	e := echo.New()

	server := &Server{
		Port: config.Port,
		Echo: e,
	}

	return server, nil
}

func (server *Server) Start() {
	address := fmt.Sprintf(":%v", server.Port)
	server.Echo.Start(address)
}
