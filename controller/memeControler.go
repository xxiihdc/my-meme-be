package controller

import (
	"fmt"
	"my-meme/adapter"
	"my-meme/env"
	"my-meme/response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func TestApi(ctx *gin.Context) {
	b, err := os.ReadFile(env.ClientSecretFile)
	if err != nil {
		fmt.Print(err)
	}

	// config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	config, err := google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope)
	if err != nil {
		fmt.Print(err)
	}

	client := adapter.GetClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		fmt.Print(err)
	}

	spreadsheetId := env.MemeTable
	readRange := "A2:J4"

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		fmt.Print(err)
	}
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
		ctx.JSON(http.StatusOK, "No data found")
	} else {
		webResponse := response.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   resp.Values,
		}
		ctx.JSON(http.StatusOK, webResponse)
	}
}

func ShowImage(ctx *gin.Context) {
	id := ctx.Param("id")
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   fmt.Sprintf("https://drive.google.com/uc?id=%s", id),
	}
	ctx.JSON(http.StatusOK, webResponse)
}
