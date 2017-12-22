package controllers

import (
//"github.com/astaxie/beego"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.Data["Website"] = "beego.mes"
	this.Data["Email"] = "astaxie@gmail.com"

	this.display()
}

func (this *IndexController) Index() {
	this.Data["Website"] = "beego.mes"
	this.Data["Email"] = "astaxie@gmail.com"

	this.display()
}
