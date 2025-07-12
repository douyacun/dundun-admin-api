package bootstrap

import (
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/middleware"
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/log"
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/routes"
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/config"
)

func RunServer() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "启动服务",
		Action: func(c *cli.Context) error {
			Start()
			return nil
		},
	}
}

func Start() {
	// gin日志/recover使用zap接收
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := gin.New()
	middleware.NewMonitor(app)

	// 代理设置
	_ = app.SetTrustedProxies([]string{"127.0.0.1"})
	app.Use(otelgin.Middleware(config.App.AppName()))
	// 注册路由
	routes.Bind(app)

	// 优雅重启
	srv := &http.Server{
		Addr:         ":" + config.App.Port(),
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      app,
	}
	gin.SetMode(gin.ReleaseMode)
	log.Infof("启动端口: %s", config.App.Port())
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// 等待中断信号
	select {
	case err := <-srvErr:
		// 启动错误
		if !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("start server err: %v", err)
		}
		return
	case <-ctx.Done():
		// 等待信号：CTRL+C.
		log.Infof("关闭服务器 ...")
		stop()
	}

	// 优雅结束服务，调用以后
	if config.App.IsProd() {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Errorf("服务器关闭: %s", err)
			return
		}
		cancel()
		time.Sleep(500 * time.Millisecond)
	}

	log.Infof("服务器进程结束.")
}
