package db

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/script/v1"
)

type DB struct {
	Service *script.Service
}

func Connect(ctx *gin.Context) (*script.Service, error) {
	return script.NewService(ctx, option.WithHTTPClient(getClient(script.SpreadsheetsScope)))
}

func getClient(scope string) *http.Client {
	b, err := os.ReadFile(os.Getenv("CLIENT_SECRET_FILE"))
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
