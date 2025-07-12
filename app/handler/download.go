package handler

import (
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/log"
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/config"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strings"
)

func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")
	dir := c.Param("dir")

	rootPath := config.App.RootPath()

	// 是否支持webp
	if supportsWebPByUA(c, c.Request) {
		ext := path.Ext(filename)
		filename = strings.Replace(filename, ext, ".webp", 1)
	}
	filepath := path.Join(rootPath, dir, filename)
	
	http.ServeFile(c.Writer, c.Request, filepath)
}

func supportsWebPByUA(ctx context.Context, r *http.Request) bool {
	ua := r.UserAgent()
	accept := r.Header.Get("Accept")
	log.CtxInfof(ctx, "ua: %s accpet %s", ua, accept)

	// 简单判断是否是 Chrome 或 Android WebView 等支持 WebP 的浏览器
	if strings.Contains(ua, "Chrome") || strings.Contains(ua, "Android") {
		return true
	}
	if strings.Contains(accept, "image/webp") {
		return true
	}
	return false
}
