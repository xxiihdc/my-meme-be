package main

import (
	"my-meme/routes"
	"net/http"

	"github.com/rs/zerolog/log"
)

// func loadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func main() {
	// loadEnv()
	log.Info().Msg("Started server!")
	router := routes.SetupRoutes()

	http.ListenAndServe(":8080", router)
}
