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

type AppCronController struct {
	BaseController
}

func (this *AppCronController) Index() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	searchType := this.GetString("search_type", "")
	searchWord := this.GetString("search_word", "")
	filters := make([]interface{}, 0)

	if searchType != "" {
		if strings.EqualFold(searchType, "name") {
			searchType2 := fmt.Sprintf("%s__icontains", searchType)
			filters = append(filters, searchType2, searchWord)
		} else {
			filters = append(filters, searchType, searchWord)
		}
	}

	result, count := models.CronGetList(page, this.pageSize, filters...)

	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})

		row["Id"] = v.Id
		row["ItemId"] = v.ItemId
		row["Name"] = v.Name
		row["ExecNum"] = v.ExecNum
		row["PrevTime"] = beego.Date(time.Unix(v.PrevTime, 0), "Y-m-d H:i:s")

		row["Status"] = v.Status
		row["UpdateTime"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["CreateTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")

		list[k] = row
	}

	this.Data["search_type"] = searchType
	this.Data["search_word"] = searchWord
	this.Data["list"] = list
	this.Data["pageLink"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("AppCronController.Index"), true).ToString()
	this.display()
}

func (this *AppCronController) Add() {

	data := new(models.AppCron)
	id, err := this.GetInt("id")

	if err == nil {
		data, _ = models.CronGetById(id)
	}

	if this.isPost() {

		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")

		data.Name = vars["name"]
		data.Desc = vars["desc"]
		data.CronSpec = vars["cron_spec"]
		data.Cmd = vars["cmd"]
		data.Timeout, _ = strconv.Atoi(vars["timeout"])
		data.Notify, _ = strconv.Atoi(vars["notify"])
		data.ItemId, _ = strconv.Atoi(vars["item_id"])

		if id > 0 {

			data.UpdateTime = time.Now().Unix()
			err := data.Update()
			if err == nil {
				msg := fmt.Sprintf("更新Cron的ID:%d|%s", id, data)
				this.uLog(msg)
				this.redirect(beego.URLFor("AppCronController.Index"))
			}
		} else {

			data.Status = 0
			data.UpdateTime = time.Now().Unix()
			data.CreateTime = time.Now().Unix()

			id, err := orm.NewOrm().Insert(data)
			if err == nil {
				msg := fmt.Sprintf("添加Cron的ID:%d", id)
				this.uLog(msg)
				this.redirect(beego.URLFor("AppCronController.Index"))
			}
		}
	}

	this.Data["data"] = data
	this.Data["id"] = this.GetString("id")

	this.display()
}

func (this *AppCronController) Lock() {

	id, err := this.GetInt("id")
	if err == nil {
		data, _ := models.CronGetById(id)

		if data.Status > 0 {
			data.Status = -1
			this.uLog("Cron锁定成功")
		} else {
			data.Status = 1
			this.uLog("Cron解锁成功")
		}
		err = data.Update()

		if err == nil {
			this.retOk("修改成功")
		}
	}
	this.retFail("修改失败")
}

func (this *AppCronController) Del() {

	id, err := this.GetInt("id")
	if err == nil {
		num, err := models.CronDelById(id)
		if err == nil {
			msg := fmt.Sprintf("删除Cron:%s,成功", num)
			this.uLog(msg)
			this.retOk(msg)
		}
	}
	this.retFail("非法参数")
}
