package cmd

import (
	"github.com/urfave/cli"

	"github.com/astaxie/beego"
	"github.com/midoks/webcron/app/libs"
	_ "github.com/midoks/webcron/app/routers"
	"github.com/midoks/webcron/app/task"
)

var Service = cli.Command{
	Name:        "service",
	Usage:       "This command starts all services",
	Description: `Start DHT services`,
	Action:      runAllService,
	Flags: []cli.Flag{
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func runAllService(c *cli.Context) error {
	libs.Init()
	task.Init()

	beego.Run()

	return nil
}
