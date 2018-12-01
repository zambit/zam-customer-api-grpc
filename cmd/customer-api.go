package main

import (
	"log"

	"git.zam.io/microservices/customer-api/service"
)

func main() {
	app, err := service.New()
	if err != nil {
		log.Fatal()
	}
	app.Run()
}
