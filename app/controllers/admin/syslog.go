package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/webcron/app/lib"
	"github.com/midoks/webcron/app/models"
	_ "runtime"
	_ "strconv"
	"strings"
	"time"
)

type SysLogController struct {
	BaseController
}

func (this *SysLogController) Index() {

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

	result, count := models.LogGetList(page, this.pageSize, filters...)

	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})

		row["id"] = v.Id
		row["uid"] = v.Uid
		row["username"] = "t"
		row["type"] = v.Type
		row["msg"] = v.Msg
		row["add_time"] = beego.Date(time.Unix(v.AddTime, 0), "Y-m-d H:i:s")

		list[k] = row
	}

	this.Data["search_type"] = searchType
	this.Data["search_word"] = searchWord
	this.Data["list"] = list
	this.Data["pageLink"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("SysLogController.Index"), true).ToString()
	this.display()
}
