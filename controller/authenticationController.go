package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"my-meme/response"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthenticationController struct{}

func NewAuthenticationController() *AuthenticationController {
	return &AuthenticationController{}
}

func (ac *AuthenticationController) Create(ctx *gin.Context) {
	userName := "Nguyen Kieu Khanh"
	expirationTime := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userName,
		"exp":  expirationTime.Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		fmt.Println("debug")
		fmt.Println(err)
	}
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tokenString,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (ac *AuthenticationController) GenerateNewGoogleToken(ctx *gin.Context) {
	b, _ := os.ReadFile(os.Getenv("CLIENT_SECRET_FILE"))
	config, _ := google.ConfigFromJSON(b,
		"https://www.googleapis.com/auth/spreadsheets",
		"https://www.googleapis.com/auth/drive",
		"https://www.googleapis.com/auth/script.deployments",
		"https://www.googleapis.com/auth/script.metrics",
		"https://www.googleapis.com/auth/script.processes",
		"https://www.googleapis.com/auth/script.projects",
	)

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	ctx.JSON(http.StatusOK, response.RespData(200, "OK", authURL))
}

func (ac *AuthenticationController) SaveToFile(ctx *gin.Context) {
	code := ctx.Query("code")
	saveToken("./token.json", code)
	ctx.JSON(200, gin.H{
		"message": "Success",
	})
}

func saveToken(path string, code string) {
	os.Remove(path)
	b, _ := os.ReadFile(os.Getenv("CLIENT_SECRET_FILE"))
	config, _ := google.ConfigFromJSON(b,
		"https://www.googleapis.com/auth/spreadsheets",
		"https://www.googleapis.com/auth/drive",
		"https://www.googleapis.com/auth/script.deployments",
		"https://www.googleapis.com/auth/script.metrics",
		"https://www.googleapis.com/auth/script.processes",
		"https://www.googleapis.com/auth/script.projects",
	)
	token, _ := config.Exchange(context.TODO(), code)
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
