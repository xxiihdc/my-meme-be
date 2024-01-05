package db

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/script/v1"
	"google.golang.org/api/sheets/v4"
)

type DB struct {
	Script *script.Service
	Sheet  *sheets.Service
}

func Connect(ctx *gin.Context) (*script.Service, error) {
	return script.NewService(ctx, option.WithHTTPClient(getClient(script.SpreadsheetsScope)))
}

func ConnectSheet(ctx *gin.Context) (*sheets.Service, error) {
	return sheets.NewService(ctx, option.WithHTTPClient(getClient(sheets.SpreadsheetsReadonlyScope)))
}

func (d *DB) FindAll(obj interface{}) ([][]interface{}, error) {
	objType := getObjType(obj)
	spreadsheetId := os.Getenv(objType)
	readRange := buildReadRange(objType)
	val, err := d.Sheet.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	return val.Values, err
}

func buildReadRange(objType string) string {
	return objType + "!" + "A:" + ""
}

func getObjType(obj interface{}) string {
	return strings.ToLower(reflect.TypeOf(obj).Name())
}

func JSONToModel(jsonData []byte, T interface{}) ([]interface{}, error) {
	return nil, nil
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
