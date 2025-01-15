package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/douyacun/go-websocket-protobuf-ts/app/log"
)

func Bind(api *gin.Engine) {
	api.GET("/gretty", func(ctx *gin.Context) {
		log.CtxInfof(ctx, "hello dundun-admin")
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "hello dundun-admin"})
	})
}
