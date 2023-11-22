package route

import (
	"tinder/config"
	"tinder/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupSwaggerRouter(config *config.Config) *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.Host = config.Server.Host + ":" + config.Server.Port
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
