package main

import (
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/bootstrap"
	"os"

	client "github.com/urfave/cli/v2"

	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/command"
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/log"
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/config"
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
		//cli.InitDB()
		// 初始化redis
		//cli.InitRDB()
		return nil
	}
	app.Commands = []*client.Command{
		bootstrap.RunServer(),
		//cronjob.Run(),
	}
	app.Commands = append(app.Commands, command.Commands...)
	if err := app.Run(os.Args); err != nil {
		log.Errorf("Run err %v", err)
	}
}
