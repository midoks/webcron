package controllers

import (
	_ "github.com/midoks/webcron/app/models"
)

type BaseController struct {
	CommonController
}

func (this *BaseController) Prepare() {

	this.initData()
	this.auth()
}
