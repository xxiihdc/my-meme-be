package routes

import (
	"fmt"
	"my-meme/controller"
	"my-meme/orm/db"
	"my-meme/repository"
	"my-meme/service"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	app := NewApp()

	router := gin.Default()
	router.Use(corsMiddleware())
	router.Use(SetJSONContentTypeMiddleware())
	apiV1Router := router.Group("/api/v1")
	{
		memeRouter := apiV1Router.Group("/meme")
		memeRouter.GET("/index", func(ctx *gin.Context) {
			app.MemeController.Index(ctx)
		})

		apiV1Router.GET("/testa", func(ctx *gin.Context) {
			app.MemeController.TestService(ctx)
		})

		apiV1Router.GET("/search", func(ctx *gin.Context) {
			app.MemeController.Index(ctx)
		})

		apiV1Router.POST("/authentication", func(ctx *gin.Context) {
			app.AuthenticationController.Create(ctx)
		})

		apiV1Router.GET("regengoogletoken", func(ctx *gin.Context) {
			app.AuthenticationController.GenerateNewGoogleToken(ctx)
		})

		apiV1Router.GET("/googletoken", func(ctx *gin.Context) {
			app.AuthenticationController.SaveToFile(ctx)
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

func authenticateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Token hợp lệ, cho phép tiếp tục xử lý request
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}

type App struct {
	MemeController *controller.MemeController

	MemeRepository *repository.MemeRepository

	MemeService *service.MemeService

	AuthenticationController *controller.AuthenticationController
}

func NewApp() *App {
	context, err := gin.CreateTestContext(nil)

	if err != nil {
		fmt.Println("118: router")
		fmt.Println(err)
		fmt.Println("==================")
	}

	script, err1 := db.Connect(context)

	sheet, _ := db.ConnectSheet(context)

	if err1 != nil {
		fmt.Println("123: router")
		fmt.Println(err1)
		fmt.Println("==================")
	}

	var db_connect db.DB
	db_connect.Script = script
	db_connect.Sheet = sheet

	memeRepo := repository.NewMemeRepositoryImpl(db_connect)

	memeService := service.NewMemeServiceImpl(memeRepo)

	memeController := controller.NewMemeController(memeService)

	authnController := controller.NewAuthenticationController()

	return &App{
		MemeController:           memeController,
		MemeService:              &memeService,
		MemeRepository:           &memeRepo,
		AuthenticationController: authnController,
	}
}
