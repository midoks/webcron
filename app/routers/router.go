package routers

import (
	"github.com/astaxie/beego"
	"github.com/midoks/webcron/app/controllers"
	"github.com/midoks/webcron/app/models"
)

func init() {

	models.Init()

	//基础
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.AutoRouter(&controllers.LoginController{})
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.SysUserController{})
	beego.AutoRouter(&controllers.SysFuncController{})
	beego.AutoRouter(&controllers.SysRoleController{})
	beego.AutoRouter(&controllers.SysLogController{})

	//后台功能开发
	beego.AutoRouter(&controllers.AppItemController{})
	beego.AutoRouter(&controllers.AppServerController{})
	beego.AutoRouter(&controllers.AppCronController{})
	beego.AutoRouter(&controllers.AppCronLogController{})
	beego.AutoRouter(&controllers.AppDebugController{})

	//前台接口
	ns := beego.NewNamespace("/v1", beego.NSAutoRouter(&controllers.ItemController{}))
	beego.AddNamespace(ns)
}
