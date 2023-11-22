package main

import (
	"tinder/internal/gin/middleware"
	"tinder/internal/gin/route"
	"tinder/internal/provider"
)

func main() {
	config := provider.NewConfig()

	swagger := route.SetupSwaggerRouter(config)
	swagger.Use(middleware.CORSMiddleware())
	swagger.Run(config.Swagger.Host + ":" + config.Swagger.Port)
}
