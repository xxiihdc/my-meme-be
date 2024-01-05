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
	FindAll() []model.Meme
	FindByKeyWord(ctx *gin.Context, keyword string) []model.Meme
}

type MemeServiceImpl struct {
	MemeRepository repository.MemeRepository
}

func (m *MemeServiceImpl) FindAll() []model.Meme {
	val, _ := m.MemeRepository.FindAll()
	return m.transformObj(val)
}

func (m *MemeServiceImpl) FindByKeyWord(ctx *gin.Context, keyword string) []model.Meme {
	val, _ := m.MemeRepository.FindByKeyWord(ctx, keyword)
	if val.Response != nil {
		result := val.Response
		return m.transformObj(m.convertJSONTo2D(result))
	} else {
		fmt.Println("No result returned from the script.")
	}
	return []model.Meme{}
}

func NewMemeServiceImpl(memeRepo repository.MemeRepository) MemeService {
	return &MemeServiceImpl{MemeRepository: memeRepo}
}

func (m *MemeServiceImpl) convertJSONTo2D(data []byte) [][]interface{} {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil
	}

	var result [][]interface{}
	for _, v := range jsonData["result"].([]interface{}) {
		if subSlice, ok := v.([]interface{}); ok {
			result = append(result, subSlice)
		} else {
			return nil
		}
	}
	return result
}

func (m *MemeServiceImpl) transformObj(arr [][]interface{}) []model.Meme {
	if len(arr) == 0 {
		return []model.Meme{}
	}
	var objects []model.Meme
	for _, subSlice := range arr {
		var id uint64
		var err interface{}
		switch subSlice[0].(type) {
		case float64:
			id = uint64(subSlice[0].(float64))
		case string:
			id, err = strconv.ParseUint(subSlice[0].(string), 10, 0)
		}

		if err != nil {

		}
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
	if objects[0].Name == "Name" {
		return objects[1:]
	}
	return objects
}
