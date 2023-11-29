package main

import (
	"fmt"
	"my-meme/routes"
	"net/http"
	"os"

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
	fmt.Println(os.Getenv("MEME_TABLE"))
	fmt.Println(os.Getenv("CLIENT_SECRET_FILE"))
	log.Info().Msg("Started server!")
	router := routes.SetupRoutes()
	http.ListenAndServe(":8888", router)
}
