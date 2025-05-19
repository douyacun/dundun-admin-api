package cronjob

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/urfave/cli/v2"
)

func Run() *cli.Command {
	return &cli.Command{
		Name:    "cronjob",
		Aliases: []string{"c"},
		Usage:   "定时任务",
		Action: func(c *cli.Context) error {
			Cronjob()
			return nil
		},
	}
}

func Cronjob() {
	c := cron.New()
	c.AddFunc("@every 1s", func() { fmt.Println("Every 1s") })
	c.Start()
}
