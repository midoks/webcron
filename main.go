package main

import (
	"github.com/astaxie/beego"
	"github.com/midoks/webcron/app/libs"
	_ "github.com/midoks/webcron/app/routers"
	"github.com/midoks/webcron/app/task"
)

func main() {

	libs.Init()
	task.Init()

	beego.Run()
}
