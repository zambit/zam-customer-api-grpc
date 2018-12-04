package main

import (
	"log"

	"git.zam.io/microservices/customer-api/service"
)

var (
	commitSHA   = "undefined"
	commitRef   = "undefined"
	commitRep   = "undefined"
	commitEnv   = "undefined"
	commitPipID = "undefined"
)

func main() {
	app, err := service.New()
	if err != nil {
		log.Fatal()
	}
	app.Run()
}
