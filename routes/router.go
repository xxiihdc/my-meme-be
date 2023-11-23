package routes

import (
	"fmt"
	"my-meme/controller"
	"my-meme/repository"
	"my-meme/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	app := NewApp()

	router := gin.Default()
	router.Use(corsMiddleware())
	router.Use(SetJSONContentTypeMiddleware())
	apiV1Router := router.Group("/api/v1")
	{
		apiV1Router.GET("/test", func(ctx *gin.Context) {
			app.MemeController.TestApi(ctx)
		})
		apiV1Router.GET("/testa", func(ctx *gin.Context) {
			app.MemeController.TestService(ctx)
		})
		apiV1Router.GET("search", func(ctx *gin.Context) {
			app.MemeController.Search(ctx)
		})
	}

	return router
}

func SetJSONContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}

func corsMiddleware() gin.HandlerFunc {
	fmt.Println("Duc test")
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

type App struct {
	MemeController *controller.MemeController

	MemeRepository *repository.MemeRepository

	MemeService *service.MemeService
}

func NewApp() *App {
	memeRepo := repository.NewMemeRepositoryImpl()

	memeService := service.NewMemeServiceImpl(memeRepo)

	memeController := controller.NewMemeController(memeService)

	return &App{
		MemeController: memeController,
		MemeService:    &memeService,
		MemeRepository: &memeRepo,
	}
}
