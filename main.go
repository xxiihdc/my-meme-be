package main

import (
	"fmt"
	"my-meme/orm/db"
	"my-meme/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var db_connect db.DB

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error on load env")
	}
}

func main() {
	loadEnv()
	log.Info().Msg("Started server!")
	context, _ := gin.CreateTestContext(nil)
	conn, err := db.Connect(context)

	db_connect.Service = conn
	if err != nil {
		fmt.Println("Have an error")
		fmt.Println(err)
		return
	}

	router := routes.SetupRoutes()
	http.ListenAndServe(":8080", router)
}
