package routes

import (
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/handler"
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"

	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/log"
)

func Bind(api *gin.Engine) {
	api.GET("/gretty", func(ctx *gin.Context) {
		log.CtxInfof(ctx, "hello dundun-admin")
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "hello word"})
	})

	dsp := api.Group("/dsp/materials/v1", middleware.RateLimitMiddleware())
	{
		dsp.GET("/:dir/:filename", handler.DownloadFile)
	}
}
