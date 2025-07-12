package command

import (
	"github.com/urfave/cli/v2"

	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/log"
)

var Commands = []*cli.Command{
	{
		Name: "gretty",
		Action: func(ctx *cli.Context) error {
			log.Infof("hello dundun-admin")
			return nil
		},
	},
}
