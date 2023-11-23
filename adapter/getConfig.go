package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"my-meme/env"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/script/v1"
	"google.golang.org/api/sheets/v4"
)

func GetSheetService(ctx *gin.Context) (*sheets.Service, error) {
	return sheets.NewService(ctx, option.WithHTTPClient(getClient(sheets.SpreadsheetsReadonlyScope)))
}

func GetScriptService(ctx *gin.Context) (*script.Service, error) {
	return script.NewService(ctx, option.WithHTTPClient(getClient(script.SpreadsheetsScope)))
}

func getClient(scope string) *http.Client {
	b, err := os.ReadFile(env.ClientSecretFile)
	if err != nil {
		fmt.Print(err)
	}

	config, err := google.ConfigFromJSON(b, scope)
	if err != nil {
		fmt.Print(err)
	}
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
