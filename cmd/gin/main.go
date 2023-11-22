package main

import (
	"tinder/internal/gin/route"
	"tinder/internal/provider"
)

func main() {

	config := provider.NewConfig()

	router := route.SetupRouter()
	router.Run(config.Server.Host + ":" + config.Server.Port)
}
