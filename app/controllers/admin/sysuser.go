package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/midoks/webcron/app/libs"
	"github.com/midoks/webcron/app/models"
	"strconv"
	"strings"
	"time"
)

type SysUserController struct {
	BaseController
}

func (this *SysUserController) Index() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	searchType := this.GetString("search_type", "")
	searchWord := this.GetString("search_word", "")
	filters := make([]interface{}, 0)

	if searchType != "" {
		if strings.EqualFold(searchType, "msg") {
			searchType2 := fmt.Sprintf("%s__icontains", searchType)
			filters = append(filters, searchType2, searchWord)
		} else {
			filters = append(filters, searchType, searchWord)
		}
	}

	result, count := models.UserGetList(page, this.pageSize, filters...)

	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})

		roleData, _ := models.RoleGetById(v.Roleid)

		row["id"] = v.Id
		row["username"] = v.Username
		row["nick"] = v.Nick
		row["sex"] = v.Sex
		row["mail"] = v.Mail
		row["tel"] = v.Tel
		row["roleid"] = v.Roleid
		if roleData.Name == "" {
			row["role_name"] = "无权限"
		} else {
			row["role_name"] = roleData.Name
		}

		row["status"] = v.Status
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")

		list[k] = row
	}

	this.Data["search_type"] = searchType
	this.Data["search_word"] = searchWord
	this.Data["list"] = list
	this.Data["pageLink"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("SysUserController.Index"), true).ToString()
	this.display()
}

func (this *SysUserController) Repwd() {

	if this.isPost() {
		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")
		tmpUser := this.user

		if vars["password"] != "" {
			tmpUser.Password = libs.Md5([]byte(vars["password"]))
		}

		tmpUser.Nick = vars["nick"]
		tmpUser.Mail = vars["mail"]
		tmpUser.Tel = vars["tel"]
		tmpUser.Roleid, _ = strconv.Atoi(vars["roleid"])
		tmpUser.Sex, _ = strconv.Atoi(vars["sex"])
		err := tmpUser.Update()
		if err == nil {
			msg := fmt.Sprintf("修改信息:%s", tmpUser)
			this.uLog(msg)
			this.redirect(beego.URLFor("SysUserController.Index"))
		}
	}
	this.display()
}

func (this *SysUserController) Add() {

	data := new(models.SysUser)
	id, err := this.GetInt("id")
	if err == nil {
		data, _ = models.UserGetById(id)
	}

	if this.isPost() {

		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")

		if id > 0 {
			data.Username = vars["username"]
			if vars["password"] != "" {
				data.Password = libs.Md5([]byte(vars["password"]))
			}

			data.Nick = vars["nick"]
			data.Mail = vars["mail"]
			data.Tel = vars["tel"]
			data.Roleid, _ = strconv.Atoi(vars["roleid"])
			data.Sex, _ = strconv.Atoi(vars["sex"])
			err := data.Update()
			if err == nil {
				msg := fmt.Sprintf("更新用户的ID:%d|%s", id, data)
				this.uLog(msg)
				this.redirect(beego.URLFor("SysUserController.Index"))
			}

		} else {

			var u models.SysUser

			u.Username = vars["username"]
			u.Password = libs.Md5([]byte(vars["password"]))
			u.Nick = vars["nick"]
			u.Mail = vars["mail"]
			u.Tel = vars["tel"]
			u.Roleid, _ = strconv.Atoi(vars["roleid"])
			u.Sex, _ = strconv.Atoi(vars["sex"])
			u.Status = 0
			u.CreateTime = time.Now().Unix()
			u.UpdateTime = time.Now().Unix()

			id, err := orm.NewOrm().Insert(&u)
			if err == nil {
				msg := fmt.Sprintf("添加用户的ID:%d", id)
				this.uLog(msg)
				this.redirect(beego.URLFor("SysUserController.Index"))
			}

		}
	}

	this.Data["data"] = data
	this.Data["id"] = this.GetString("id")

	roleList, _ := models.RoleGetAll()
	this.Data["roleList"] = roleList

	this.display()
}

func (this *SysUserController) Lock() {

	id, err := this.GetInt("id")

	if err == nil {
		data, _ := models.UserGetById(id)
		data.Password = ""
		this.retResult(1, "123123", data)
	}

	fmt.Println(id, err)
	this.retResult(1, "123123")
}
