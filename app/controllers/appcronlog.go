package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	"github.com/midoks/webcron/app/libs"
	"github.com/midoks/webcron/app/models"
	// "strconv"
	"strings"
	"time"
)

type AppCronLogController struct {
	BaseController
}

func (this *AppCronLogController) Index() {
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

	result, count := models.CronLogGetList(page, this.pageSize, filters...)

	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})

		row["Id"] = v.Id
		row["CronId"] = v.CronId
		row["Output"] = v.Output
		row["Error"] = v.Error

		row["ProcessTime"] = v.ProcessTime

		row["Status"] = v.Status
		row["CreateTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")

		list[k] = row
	}

	this.Data["search_type"] = searchType
	this.Data["search_word"] = searchWord
	this.Data["list"] = list
	this.Data["pageLink"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("AppCronLogController.Index"), true).ToString()
	this.display()
}
