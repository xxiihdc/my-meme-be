package controller

import (
	"my-meme/response"
	"my-meme/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemeController struct {
	memeService service.MemeService
}

func NewMemeController(memeService service.MemeService) *MemeController {
	return &MemeController{
		memeService: memeService,
	}
}

func (mc *MemeController) Index(ctx *gin.Context) {
	keyword := ctx.Query("q")
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
	}

	if keyword == "" {
		webResponse.Data = mc.memeService.FindAll()
	} else {
		webResponse.Data = mc.memeService.FindByKeyWord(ctx, keyword)
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (mc *MemeController) TestService(ctx *gin.Context) {
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   mc.memeService.FindAll(),
	}
	ctx.JSON(http.StatusOK, webResponse)
}
