package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

func GetClient(config *oauth2.Config) *http.Client {
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
