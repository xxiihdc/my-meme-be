package repository

import (
	"my-meme/model"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/script/v1"
)

type MemeRepository interface {
	Repository[model.Meme]
	FindByKeyWord(ctx *gin.Context, key string) (*script.Operation, error)
}
