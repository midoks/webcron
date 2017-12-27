package controllers

import (
	"fmt"
	// "github.com/astaxie/beego"
	// "github.com/midoks/webcron/app/lib"
	"github.com/midoks/webcron/app/models"
	"strings"
)

type SysFuncController struct {
	BaseController
}

func (this *SysFuncController) Index() {

	result := models.FuncGetList()

	//对(栏目名)填充内容,利于后台观看
	for i := 0; i < len(result); i++ {
		for ci := 0; ci < len(result[i].List); ci++ {
			//println(result[i].List[ci].Name, len(result[i].List[ci].Name))
			fillcount := 20 - len(result[i].List[ci].Name)
			if fillcount > 0 && len(result[i].List[ci].Name) < 16 {
				tmp := strings.Repeat(" ", 20-len(result[i].List[ci].Name))
				result[i].List[ci].Name = fmt.Sprintf("%s%s", result[i].List[ci].Name, tmp)
			}
		}
	}

	this.Data["list"] = result
	this.display()
}
