package routers

import (
	"github.com/astaxie/beego"
	"github.com/midoks/webcron/app/controllers/admin"
	"github.com/midoks/webcron/app/models"
)

func init() {

	models.Init()

	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.AutoRouter(&controllers.LoginController{})
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.SysUserController{})
	beego.AutoRouter(&controllers.SysFuncController{})
	beego.AutoRouter(&controllers.SysRoleController{})
	beego.AutoRouter(&controllers.SysLogController{})

}
