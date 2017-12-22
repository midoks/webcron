package controllers

import (
//"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.mes"
	this.Data["Email"] = "astaxie@gmail.com"

	this.display()
}
