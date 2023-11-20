package routes

import (
	"my-meme/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	apiV1Router := router.Group("/api/v1")
	{
		apiV1Router.GET("/test", func(ctx *gin.Context) {
			controller.TestApi(ctx)
		})
		apiV1Router.GET("/test2/:id", func(ctx *gin.Context) {
			controller.ShowImage(ctx)
		})
	}

	return router
}
