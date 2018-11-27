package main

import "git.zam.io/microservices/customer-api/application"

func main() {
	app, err := application.New()
	if err != nil {
		panic(err)
	}
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
