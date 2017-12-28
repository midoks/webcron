package controllers

import ()

type BaseController struct {
	CommonController
}

func (this *BaseController) Prepare() {

	this.initData()

	this.auth()

	this.initMenuData()
}
