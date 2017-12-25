package controllers

//"github.com/astaxie/beego"

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {

	this.display()

}
