package route

import (
	"tinder/internal/gin/handler"
	"tinder/internal/gin/middleware"

	"github.com/gin-gonic/gin"
)

// @title           Tinder API
// @version         0.0.1
// @host      0.0.0.0:3002
// @description.markdown
// @schemes   http
func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CORSMiddleware())
	api := r.Group("/people")
	{
		api.POST("/add_and_match", handler.AddSinglePersonAndMatch)
		api.DELETE("/:userName", handler.RemoveSinglePerson)
		api.POST("/:userName/query_single_person/:number", handler.QuerySinglePerson)
	}

	return r
}
