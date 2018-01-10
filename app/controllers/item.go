package controllers

import (
// "fmt"
// "github.com/astaxie/beego"
// "github.com/astaxie/beego/orm"
// "github.com/midoks/webcron/app/libs"
// "github.com/midoks/webcron/app/models"
// "strconv"
// "strings"
// "time"
)

type ItemController struct {
	CommonController
}

// func (this *ItemController) Prepare() {

// 	this.initData()
// 	this.auth()
// }

func (this *ItemController) Index() {
	this.retOk("ok")
}
