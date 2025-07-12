package middleware

import (
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/log"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strconv"
)

const DefaultMetricPath = "/metrics"

/**
监测指标
1. 延迟
2. 错误
3. 请求量
4. 饱和度
*/

// https://github.com/slok/go-http-metrics/blob/master/metrics/prometheus/prometheus.go
// https://github.com/PayU/prometheus-api-metrics/tree/master

// 定义指标。会自动注入
var (
	httpLabelNames    = []string{"handler", "method", "status"}
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "HTTP request total.",
		},
		httpLabelNames,
	)

	// http.ServeFile 没办法通过中间件获取文件大小
	//httpResponseBytes = prometheus.NewCounterVec(
	//	prometheus.CounterOpts{
	//		Name: "http_response_size_bytes_total",
	//		Help: "Total number of bytes served in HTTP responses",
	//	},
	//	httpLabelNames,
	//)
)

func NewMonitor(e *gin.Engine) {
	// 注册metrics路由
	e.GET(DefaultMetricPath, prometheusHandler())
	// 注册中间件
	e.Use(HandleFunc())
}

func HandleFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		ctx.Next()
		log.CtxInfof(ctx, "response.size %.2f", float64(ctx.Writer.Size()))
		httpRequestsTotal.WithLabelValues(path, ctx.Request.Method, strconv.Itoa(ctx.Writer.Status())).Inc()
		//httpResponseBytes.WithLabelValues(path, ctx.Request.Method, strconv.Itoa(ctx.Writer.Status())).Add(float64(ctx.Writer.Size()))
	}
}
func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
