package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Port string
	Echo *echo.Echo
}

func NewServer(port string) *Server {
	// Create Echo Server
	e := echo.New()
	e.HideBanner = true
	e.Pre(middleware.RemoveTrailingSlash())

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server := &Server{
		Port: port,
		Echo: e,
	}
	return server
}

func (server *Server) HandleShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	sig := <-quit
	server.Echo.Logger.Errorf("Got signal: %v, shutting down the server\n", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Echo.Shutdown(ctx); err != nil {
		server.Echo.Logger.Fatal(err)
	}
}

func (server *Server) Start() {
	address := fmt.Sprintf(":%s", server.Port)
	server.Echo.Logger.Fatal(server.Echo.Start(address))
}
