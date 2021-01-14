package main

import (
	"log"

	"app/pkg/app"
)

func main() {
	log.Println("golang-microservice-template")

	app1 := app.NewApp()
	app1.Start()
}
