package main

import (
	"context"
	"encoding/json"
	"fmt"
	"my-meme/response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func getClient(config *oauth2.Config) *http.Client {
	tok, err := tokenFromFile("token.json")
	if err != nil {
		fmt.Print("===========ERROR========")
	}
	return config.Client(context.Background(), tok)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func main() {
	log.Info().Msg("Started server!")
	router := gin.Default()
	router.GET("ok", func(ctx *gin.Context) {
		webResponse := response.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   "OK",
		}
		ctx.JSON(http.StatusOK, webResponse)
	})

	router.GET("/test/sheets", func(ctx *gin.Context) {
		b, err := os.ReadFile("client_secret_473522429728-ajiotkm33vahlcn9dgci5djgsdg7onkk.apps.googleusercontent.com.json")
		if err != nil {
			fmt.Print(err)
		}

		config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
		if err != nil {
			fmt.Print(err)
		}

		client := getClient(config)

		srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			fmt.Print(err)
		}

		spreadsheetId := "14XPDaQ6eU6lMycEoEpqaPAplZJBdsWmQzUvvKG36Ik8"
		readRange := "A2:J4"

		resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
		if err != nil {
			fmt.Print(err)
		}

		webResponse := response.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   resp,
		}
		ctx.JSON(http.StatusOK, webResponse)

		// if len(resp.Values) == 0 {
		// 	fmt.Println("No data found.")
		// 	ctx.JSON(http.StatusOK, "No data found")
		// } else {
		// 	webResponse := response.Response{
		// 		Code:   http.StatusOK,
		// 		Status: "OK",
		// 		Data:   resp.Values,
		// 	}
		// 	ctx.JSON(http.StatusOK, webResponse)
		// }
	})

	router.Run()
}
