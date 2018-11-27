package application

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"git.zam.io/microservices/customer-api/pkg/config"
	"git.zam.io/microservices/customer-api/service/endpoints"
	"git.zam.io/microservices/customer-api/service/sd"
	"git.zam.io/microservices/customer-api/service/service"
	httptransport "git.zam.io/microservices/customer-api/service/transport/http"
	"github.com/go-kit/kit/sd/consul"
)

type Application struct {
	server *http.Server
	sd     *consul.Registrar
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
	reg := sd.ConsulRegister()
	app.sd = reg.(*consul.Registrar)
	app.sd.Register()
	return nil
}

func (app *Application) Run() error {
	h := httptransport.NewHTTPHandler(endpoints.MakeServerEndpoints(&service.CustomerAPIService{}))

	go func() {
		log.Fatal(http.ListenAndServe(config.Config().GetString("application.host")+":"+config.Config().GetString("application.http.port"), h))
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	app.shutdown()

	return nil
}

func (app *Application) shutdown() {
	app.sd.Deregister()
}
