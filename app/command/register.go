package command

import (
	"github.com/urfave/cli/v2"

	"github.com/douyacun/go-websocket-protobuf-ts/app/log"
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
