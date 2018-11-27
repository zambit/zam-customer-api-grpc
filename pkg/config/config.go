package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var currentConfig *viper.Viper
var onceInit sync.Once

func Config() *viper.Viper {
	onceInit.Do(func() {
		currentConfig = viper.New()
		currentConfig.SetEnvPrefix("CUSTOMER_API")
		currentConfig.AutomaticEnv()
		currentConfig.SetConfigType("yaml")
		currentConfig.AddConfigPath("./etc/")

		runmode := currentConfig.Get("RUNMODE").(string)
		if len(runmode) == 0 {
			runmode = "dev"
		}
		currentConfig.SetConfigFile("./etc/" + runmode + ".yaml")
		err := currentConfig.ReadInConfig()
		if err != nil {
			log.Fatalf("YAML load config failed: %s\n", err)
		}

		log.Println("--------------------------")
		log.Println("Config Loaded: \t\t", "OK")
		log.Println("Run mode: \t\t\t", runmode)

	})

	return currentConfig
}
