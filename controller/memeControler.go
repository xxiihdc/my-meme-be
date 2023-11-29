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
		webResponse.Data = mc.memeService.FindAll(ctx)
	} else {
		webResponse.Data = mc.memeService.FindByKeyWord(ctx, keyword)
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// func (*MemeController) TestApi(ctx *gin.Context) {
// 	srv, err := adapter.GetSheetService(ctx)
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	spreadsheetId := env.MemeTable
// 	readRange := "A2:J4"

// 	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	if len(resp.Values) == 0 {
// 		fmt.Println("No data found.")
// 		ctx.JSON(http.StatusOK, "No data found")
// 	} else {
// 		webResponse := response.Response{
// 			Code:   http.StatusOK,
// 			Status: "OK",
// 			Data:   resp.Values,
// 		}
// 		ctx.JSON(http.StatusOK, webResponse)
// 	}
// }

func (mc *MemeController) TestService(ctx *gin.Context) {
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   mc.memeService.FindAll(ctx),
	}
	ctx.JSON(http.StatusOK, webResponse)
}
