package application

import (
	"log"
	"net/http"

	"git.zam.io/microservices/customer-api/service/endpoints"
	http2 "git.zam.io/microservices/customer-api/service/http"
	"git.zam.io/microservices/customer-api/service/sd"
	"git.zam.io/microservices/customer-api/service/service"
	"github.com/go-kit/kit/sd/consul"
)

type Application struct {
	server *http.Server
	sd     consul.Registrar
}

func New() (*Application, error) {
	app := &Application{}
	err := app.init()
	if err != nil {
		log.Fatal(err)
	}
	return app, nil
}

func (app *Application) init() error {
	app.server = &http.Server{}
	registar := sd.ConsulRegister()
	registar.Register()
	return nil
}

func (app *Application) Run() error {
	h := http2.NewHTTPHandler(endpoints.MakeServerEndpoints(service.CustomerAPIService{}))

	log.Fatal(http.ListenAndServe(":3000", h))
	return nil
}
