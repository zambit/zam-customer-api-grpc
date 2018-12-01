package service

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/satori/go.uuid"

	"git.zam.io/microservices/customer-api/pkg/config"
)

func ConsulRegister() (registar sd.Registrar) {
	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Service discovery domain. In this example we use Consul.
	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		consulConfig.Address = config.Config().GetString("consul.host") + ":" + config.Config().GetString("consul.port")
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consulsd.NewClient(consulClient)
	}

	check := api.AgentServiceCheck{
		HTTP:     "http://" + config.Config().GetString("application.host") + ":" + config.Config().GetString("application.http.port") + "/health",
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Customer API health checks",
	}

	asr := api.AgentServiceRegistration{
		ID:      "customer-api-" + uuid.NewV4().String(), // unique service ID
		Name:    "customer-api",
		Address: "http://" + config.Config().GetString("application.host"),
		Port:    config.Config().GetInt("application.http.port"),
		Tags:    []string{"customer-api", "http"},
		Check:   &check,
	}
	registar = consulsd.NewRegistrar(client, &asr, logger)
	return
}
