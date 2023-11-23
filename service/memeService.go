package service

import (
	"encoding/json"
	"fmt"
	"my-meme/model"
	"my-meme/repository"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MemeService interface {
	FindAll(ctx *gin.Context) []model.Meme
	FindByKeyWord(ctx *gin.Context, keyword string) string
}

type MemeServiceImpl struct {
	MemeRepository repository.MemeRepository
}

func (m *MemeServiceImpl) FindAll(ctx *gin.Context) []model.Meme {
	val, _ := m.MemeRepository.GetAll(ctx)
	var objects []model.Meme
	for _, subSlice := range val {
		id, _ := strconv.ParseUint(subSlice[0].(string), 10, 0)
		obj := model.Meme{
			ID:          uint(id),
			DriveId:     subSlice[1].(string),
			Name:        subSlice[2].(string),
			Description: subSlice[3].(string),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		objects = append(objects, obj)
	}
	return objects[1:]
}

func (m *MemeServiceImpl) FindByKeyWord(ctx *gin.Context, keyword string) string {
	val, _ := m.MemeRepository.FindByKeyWord(ctx, keyword)
	if val.Response != nil {
		result := val.Response
		var jsonData map[string]interface{}
		json.Unmarshal([]byte(result), &jsonData)
		return jsonData["result"].(string)
	} else {
		fmt.Println("No result returned from the script.")
	}
	return "OK"
}

func NewMemeServiceImpl(memeRepo repository.MemeRepository) MemeService {
	return &MemeServiceImpl{MemeRepository: memeRepo}
}
