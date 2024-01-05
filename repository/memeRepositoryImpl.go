package repository

import (
	"fmt"
	"log"
	"my-meme/adapter"
	"my-meme/model"
	"my-meme/orm/db"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/script/v1"
)

type MemeRepositoryImpl struct {
	database db.DB
}

func (r *MemeRepositoryImpl) FindAll() ([][]interface{}, error) {
	service := r.database.Sheet
	godotenv.Load()
	spreadsheetId := os.Getenv("MEME_TABLE")

	readRange := "Sheet1!A:F"

	meme := model.Meme{}

	r.database.FindAll(meme)

	val, err := service.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		fmt.Println("duc debug")
		fmt.Println(err)
		fmt.Println("duc debug")
	}
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

	req := &script.ExecutionRequest{
		Function:   "searchTerm",
		Parameters: []interface{}{keyword},
	}

	resp, err := service.Scripts.Run(scriptID, req).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to execute script: %v", err)
	}
	return resp, err
}

func NewMemeRepositoryImpl(db db.DB) MemeRepository {
	return &MemeRepositoryImpl{
		database: db,
	}
}
