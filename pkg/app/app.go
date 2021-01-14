package app

import (
	"log"
)

type App struct {
	Config   Config
	Server   *Server
	Database *Database
	Logger   log.Logger
	// TODO: Put DB as depedencies
}

func (app *App) Start() error {
	server := app.Server

	server.Echo.Any("/", app.HealthCheck)

	server.Start()
	// TODO: Graceful Shutdown

	return nil
}

func NewApp() *App {
	config := NewConfig()
	server, _ := NewServer(config)
	database, _ := NewDatabase(config)

	app := &App{
		Config:   config,
		Server:   server,
		Database: database,
	}

	return app
}
