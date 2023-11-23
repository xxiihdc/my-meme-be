package repository

import (
	"fmt"
	"log"
	"my-meme/adapter"
	"my-meme/env"
	"my-meme/model"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/script/v1"
)

type MemeRepositoryImpl struct {
}

func (r *MemeRepositoryImpl) GetAll(ctx *gin.Context) ([][]interface{}, error) {
	service, err := adapter.GetSheetService(ctx)

	if err != nil {
		panic("Error on get client Google API")
	}

	spreadsheetId := env.MemeTable
	readRange := "Sheet1!A:F"
	val, err := service.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	return val.Values, err
}

func (r *MemeRepositoryImpl) GetByID(ctx *gin.Context, id int) (*model.Meme, error) {
	meme := new(model.Meme)
	return meme, nil
}

func (r *MemeRepositoryImpl) Create(ctx *gin.Context, t model.Meme) (*model.Meme, error) {
	meme := new(model.Meme)
	return meme, nil
}

func (r *MemeRepositoryImpl) Update(ctx *gin.Context, t model.Meme) (*model.Meme, error) {
	meme := new(model.Meme)
	return meme, nil
}

func (r *MemeRepositoryImpl) Delete(ctx *gin.Context, id int) error {
	return nil
}

func (r *MemeRepositoryImpl) FindByKeyWord(ctx *gin.Context, keyword string) (*script.Operation, error) {
	service, err := adapter.GetScriptService(ctx)

	if err != nil {
		panic("Error on get client Google API")
	}

	scriptID := "AKfycbwVrTvBNIu-K3DR8epeGoxVzQWRAR0BUia87TYBPWR7D_TTQFZQnp_ECVfG3eFSwhQNIw"

	payload := map[string]interface{}{
		"function": "searchTerm",
		"parameters": []interface{}{
			"Kieu Khanh",
		},
	}

	req := &script.ExecutionRequest{
		Function: "searchTerm",
		Parameters: []interface{}{
			scriptID,
			payload,
		},
	}

	resp, err := service.Scripts.Run(scriptID, req).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to execute script: %v", err)
	}
	data := resp.Response[0]
	fmt.Println(data)
	return resp, err
}

func NewMemeRepositoryImpl() MemeRepository {
	return &MemeRepositoryImpl{}
}
