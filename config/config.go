package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server  *ServerConfig
	Swagger *SwaggerConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type SwaggerConfig struct {
	Host string
	Port string
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		Server: &ServerConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
		},
		Swagger: &SwaggerConfig{
			Host: viper.GetString("swagger.host"),
			Port: viper.GetString("swagger.port"),
		},
	}
}
