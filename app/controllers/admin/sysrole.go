package controllers

import (
	"fmt"
	// "github.com/astaxie/beego"
	// "github.com/midoks/webcron/app/lib"
	"github.com/midoks/webcron/app/models"
	// "strings"
	// "time"
)

type SysRoleController struct {
	BaseController
}

func (this *SysRoleController) Index() {

	list, _ := models.RoleGetAll()

	this.Data["list"] = list

	this.display()
}

func (this *SysRoleController) Add() {
	this.display()

	if this.isPost() {
		fmt.Println("--post--")

	}

	id, err := this.GetInt("id")
	//if err == nil {
	data, _ := models.RoleGetById(id)
	this.Data["data"] = data
	fmt.Println(data)
	//}

	fmt.Println("--test--", err)

	funcList := models.FuncGetList()

	this.Data["funcList"] = funcList

}
