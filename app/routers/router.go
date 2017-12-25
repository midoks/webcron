package routers

import (
	"github.com/astaxie/beego"
	"github.com/midoks/webcron/app/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.AutoRouter(&controllers.IndexController{})
}
