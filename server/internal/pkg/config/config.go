package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
)

func NewConfig() *viper.Viper {
	log.Printf("Initializing env: %v\n", os.Getenv("APP_ENV"))

	envConf := os.Getenv("APP_ENV")
	if envConf == "" {
		log.Fatal("Error loading .env variables")
	}

	configFileString := ""
	if envConf == "local" {
		configFileString += "config/local.yml"
	} else if envConf == "prod" {
		configFileString += "config/prod.yml"
	} else {
		log.Fatal("Unexpected value of APP_ENV")
	}

	flag.StringVar(&envConf, "conf", configFileString, "config path")
	flag.Parse()

	log.Printf("loaded conf file: %v\n", configFileString)

	return getConfig(envConf)
}

func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
