package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/webcron/app/lib"
	"github.com/midoks/webcron/app/models"
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

	result, count := models.UserGetList(page, this.pageSize)

	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})

		row["id"] = v.Id
		row["user_name"] = v.Username
		row["nick"] = v.Nick
		row["sex"] = v.Sex
		row["mail"] = v.Mail
		row["tel"] = v.Tel
		row["roleid"] = v.Roleid
		row["status"] = v.Status
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")

		list[k] = row
	}

	this.Data["search_type"] = searchType
	this.Data["search_word"] = searchWord
	this.Data["list"] = list
	this.Data["pageLink"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("SysLogController.Index"), true).ToString()
	this.display()
}
