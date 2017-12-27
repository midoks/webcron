package main

import (
	"github.com/astaxie/beego"
	"github.com/midoks/webcron/app/libs"
	_ "github.com/midoks/webcron/app/routers"
)

func main() {

	libs.Init()

	beego.Run()

	// beego.EnableXSRF = true
	// beego.XSRFKEY = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	// beego.XSRFExpire = 3600 //过期时间，默认1小时
}
