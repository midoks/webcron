package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/midoks/webcron/app/libs"
	"github.com/midoks/webcron/app/models"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type AppServerController struct {
	BaseController
}

func (this *AppServerController) Index() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	searchType := this.GetString("search_type", "")
	searchWord := this.GetString("search_word", "")
	filters := make([]interface{}, 0)

	if searchType != "" {
		if strings.EqualFold(searchType, "desc") {
			searchType2 := fmt.Sprintf("%s__icontains", searchType)
			filters = append(filters, searchType2, searchWord)
		} else {
			filters = append(filters, searchType, searchWord)
		}
	}

	result, count := models.ServerGetList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})

		row["Id"] = v.Id
		row["Desc"] = v.Desc
		row["Ip"] = v.Ip
		row["Port"] = v.Port
		row["Type"] = v.Type
		row["User"] = v.User
		row["Pwd"] = v.Pwd

		row["Status"] = v.Status
		row["UpdateTime"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["CreateTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")

		list[k] = row
	}

	this.Data["search_type"] = searchType
	this.Data["search_word"] = searchWord
	this.Data["list"] = list
	this.Data["pageLink"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("AppServerController.Index"), true).ToString()
	this.display()
}

func (this *AppServerController) SearchAjax() {
	page := 1
	filters := make([]interface{}, 0)

	qstr := this.GetString("q")
	if qstr == "" {
		//this.retFail("搜索词不能为空")
	} else {
		q, err := strconv.Atoi(qstr)
		filters = append(filters, "status", 1)
		if err == nil {
			filters = append(filters, "id", q)
		} else {
			filters = append(filters, "desc__icontains", qstr)
		}
	}

	result, _ := models.ServerGetList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})

		row["Id"] = v.Id
		row["Desc"] = v.Desc
		row["Ip"] = v.Ip
		row["Port"] = v.Port
		row["Type"] = v.Type
		row["User"] = v.User
		row["Pwd"] = v.Pwd

		row["Status"] = v.Status
		row["UpdateTime"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["CreateTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")

		list[k] = row
	}
	this.retOk("ok", list)
}

func (this *AppServerController) Add() {

	data := new(models.AppServer)
	id, err := this.GetInt("id")

	rsaKey := ""
	// fmt.Println(beego.AppConfig.String("local.id_rsa"))
	if beego.AppConfig.String("local.id_rsa") != "" {
		// fmt.Println("local.id_rsa ok")
		rsaContent, rsaErr := ioutil.ReadFile(fmt.Sprintf("conf/%s", beego.AppConfig.String("local.id_rsa")))
		if rsaErr == nil {
			rsaKey = string(rsaContent)
			fmt.Println(rsaKey)
		}
	}
	this.Data["rsaKey"] = rsaKey

	if err == nil {
		data, _ = models.ServerGetById(id)
	}

	if this.isPost() {

		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")

		data.Desc = vars["desc"]
		data.Ip = vars["ip"]
		data.Port, _ = strconv.Atoi(vars["port"])
		data.Type, _ = strconv.Atoi(vars["type"])

		data.User = vars["user"]
		data.Pwd = vars["pwd"]

		if id > 0 {

			data.UpdateTime = time.Now().Unix()
			err := data.Update()
			if err == nil {
				msg := fmt.Sprintf("更新Server服务器的ID:%d|%s", id, data)
				this.uLog(msg)
				this.redirect(beego.URLFor("AppServerController.Index"))
			}
		} else {

			data.Status = 0
			data.UpdateTime = time.Now().Unix()
			data.CreateTime = time.Now().Unix()

			id, err := orm.NewOrm().Insert(data)
			if err == nil {
				msg := fmt.Sprintf("添加Server服务器的ID:%d", id)
				this.uLog(msg)
				this.redirect(beego.URLFor("AppServerController.Index"))
			}
		}
	}

	this.Data["data"] = data
	this.Data["id"] = this.GetString("id")

	roleList, _ := models.RoleGetAll()
	this.Data["roleList"] = roleList

	this.display()
}

func (this *AppServerController) Lock() {

	id, err := this.GetInt("id")
	if err == nil {
		data, _ := models.ServerGetById(id)

		if data.Status > 0 {
			data.Status = -1
			this.uLog("Server服务器成功")
		} else {
			data.Status = 1
			this.uLog("Server服务器成功")
		}
		err = data.Update()

		if err == nil {
			this.retOk("修改成功")
		}
	}
	this.retFail("修改失败")
}

func (this *AppServerController) Del() {

	id, err := this.GetInt("id")
	if err == nil {
		num, err := models.ServerDelById(id)
		if err == nil {
			msg := fmt.Sprintf("删除Server服务器%s成功", num)
			this.uLog(msg)
			this.retOk(msg)
		}
	}
	this.retFail("非法参数")
}
