package controllers

//"github.com/astaxie/beego"

type LoginController struct {
	BaseController
}

func (this *LoginController) Index() {

	this.display()
}
