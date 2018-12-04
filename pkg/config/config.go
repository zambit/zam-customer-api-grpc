package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/pflag"
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

		pflag.String("config", "", "Config file")
		pflag.Parse()
		err := currentConfig.BindPFlags(pflag.CommandLine)
		if err != nil {
			log.Fatal(err)
		}
		configFile := currentConfig.GetString("config")
		if len(configFile) != 0 {
			currentConfig.SetConfigFile(configFile)
		} else {
			runmode := currentConfig.Get("RUNMODE").(string)
			if len(runmode) == 0 {
				runmode = "dev"
			}
			currentConfig.SetConfigFile("./etc/" + runmode + ".yaml")
		}

		err = currentConfig.ReadInConfig()
		if err != nil {
			log.Fatalf("YAML load config failed: %s\n", err)
		}
		fmt.Printf("Config loaded from: %s\n\n", currentConfig.ConfigFileUsed())

	})

	return currentConfig
}
