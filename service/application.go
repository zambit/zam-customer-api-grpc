package service

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"git.zam.io/microservices/customer-api/pb"
	"git.zam.io/microservices/customer-api/pkg/config"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
)

type Application struct {
	server *http.Server
	sd     *consul.Registrar
	logger log.Logger
}

func New() (*Application, error) {
	app := &Application{}
	err := app.init()
	if err != nil {
		app.logger.Log(err)
	}
	return app, nil
}

func (app *Application) init() error {
	app.server = &http.Server{}
	{
		app.logger = log.NewLogfmtLogger(os.Stderr)
		app.logger = log.With(app.logger, "ts", log.DefaultTimestamp)
		app.logger = log.With(app.logger, "caller", log.DefaultCaller)
	}

	reg := ConsulRegister()
	app.sd = reg.(*consul.Registrar)
	app.sd.Register()

	return nil
}

func (app *Application) Run() {
	h := NewHTTPServer(MakeServerEndpoints(&CustomerAPIService{}, app.logger), app.logger)

	// run HTTP-server
	go func() {
		err := app.logger.Log(http.ListenAndServe(config.Config().GetString("application.host")+":"+config.Config().GetString("application.http.port"), h))
		if err != nil {
			os.Exit(1)
		}
	}()

	// run GRPC-server
	go func() {
		listener, err := net.Listen("tcp", config.Config().GetString(config.Config().GetString("application.host")+":"+config.Config().GetString("application.grpc.port")))
		if err != nil {
			app.logger.Log(err)
			os.Exit(1)
		}

		gRPCServer := grpc.NewServer()
		pb.RegisterCustomerAPIServiceGRPCServer(gRPCServer, &CustomerAPIGRPCServer{})
		err = gRPCServer.Serve(listener)
		if err != nil {
			os.Exit(1)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	<-sigs

	app.shutdown()
}

func (app *Application) shutdown() {
	app.sd.Deregister()
}
