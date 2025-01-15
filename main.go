package main

import (
	"os"

	client "github.com/urfave/cli/v2"

	"github.com/douyacun/go-websocket-protobuf-ts/app/cli"
	"github.com/douyacun/go-websocket-protobuf-ts/app/command"
	"github.com/douyacun/go-websocket-protobuf-ts/app/log"
	"github.com/douyacun/go-websocket-protobuf-ts/bootstrap"
	"github.com/douyacun/go-websocket-protobuf-ts/config"
)

func main() {
	// @doc https://cli.urfave.org/v2/getting-started/
	app := &client.App{
		Suggest:              true,
		EnableBashCompletion: true,
	}
	app.Flags = []client.Flag{
		&client.StringFlag{
			Name:        "conf",
			Usage:       "-c <配置文件>",
			Required:    false,
			Aliases:     []string{"c"},
			DefaultText: "./config.ini",
		},
	}
	app.Before = func(ctx *client.Context) error {
		// 初始化配置
		config.Init(ctx.String("conf"))
		// 初始化日志
		log.Init()
		// 初始化mysql
		cli.InitDB()
		// 初始化redis
		cli.InitRDB()
		return nil
	}
	app.Commands = []*client.Command{
		bootstrap.RunServer(),
	}
	app.Commands = append(app.Commands, command.Commands...)
	if err := app.Run(os.Args); err != nil {
		log.Errorf("Run err %v", err)
	}
}
