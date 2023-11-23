package repository

import "github.com/gin-gonic/gin"

type Repository[T interface{}] interface {
	GetAll(ctx *gin.Context) ([][]interface{}, error)
	GetByID(ctx *gin.Context, id int) (*T, error)
	Create(ctx *gin.Context, t T) (*T, error)
	Update(ctx *gin.Context, t T) (*T, error)
	Delete(ctx *gin.Context, id int) error
}
