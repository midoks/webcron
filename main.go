package main

import (
	"github.com/astaxie/beego"
	"github.com/midoks/webcron/app/libs"
	_ "github.com/midoks/webcron/app/routers"
)

func main() {

	libs.Init()

	beego.Run()

}
