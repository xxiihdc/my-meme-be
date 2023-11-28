package main

import (
	"fmt"
	"my-meme/routes"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error on load env")
	}
}

func main() {
	loadEnv()
	// memeRepo := repository.NewMemeRepositoryImpl()
	// memeService := service.NewMemeServiceImpl(memeRepo)
	log.Info().Msg("Started server!")
	router := routes.SetupRoutes()
	http.ListenAndServe(":8888", router)
}
