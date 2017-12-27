package models

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type SysFunc struct {
	Id         int
	Name       string
	Controller string
	Action     string
	Type       string
	IsMenu     int64
	Icon       string
	Desc       string
	Sort       int
	Status     int
	UpdateTime int64
	CreateTime int64
}

type SysFuncNav struct {
	Info      SysFunc
	List      []SysFunc
	MenuOpen  bool
	ListCount int
}

func (u *SysFunc) TableName() string {
	return "sys_func"
}

func (u *SysFunc) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func FuncGetNav(curController string, curAction string) (navNow []SysFuncNav, menuNameNow string, funcNameNow string) {

	o := orm.NewOrm()
	var list []SysFunc

	res, _ := o.Raw("select * from sys_func where pid=? and status=? order by sort asc", 0, 1).QueryRows(&list)
	nav := make([]SysFuncNav, len(list))
	var curMenuName string = ""
	var curMenuFuncName string = ""

	if res > 0 {
		for i := 0; i < len(list); i++ {
			var cList []SysFunc
			cres, _ := o.Raw("select * from sys_func where pid=? and status=? order by sort asc", list[i].Id, 1).QueryRows(&cList)
			if cres > 0 {
				nav[i].Info = list[i]
				nav[i].List = cList
				nav[i].ListCount = len(cList)
				nav[i].MenuOpen = false

				for ci := 0; ci < len(cList); ci++ {
					if strings.EqualFold(cList[ci].Controller, curController) && strings.EqualFold(cList[ci].Action, curAction) {
						nav[i].MenuOpen = true
						// fmt.Println("debug:", cList[ci].Controller, curController, cList[ci].Action, curAction)
						curMenuName = list[i].Name
						curMenuFuncName = cList[ci].Name
					}
				}
			}
		}
	}

	return nav, curMenuName, curMenuFuncName
}

func FuncGetList() []SysFuncNav {

	o := orm.NewOrm()
	var list []SysFunc

	res, _ := o.Raw("select * from sys_func where pid=? and status=? order by sort asc", 0, 1).QueryRows(&list)
	nav := make([]SysFuncNav, len(list))

	if res > 0 {
		for i := 0; i < len(list); i++ {
			var cList []SysFunc
			cres, _ := o.Raw("select * from sys_func where pid=? and status=? order by sort asc", list[i].Id, 1).QueryRows(&cList)
			if cres > 0 {
				nav[i].Info = list[i]
				nav[i].List = cList
				nav[i].ListCount = len(cList)
			}
		}
	}

	return nav
}

func FuncGetById(id int) (*SysFunc, error) {

	sysfunc := new(SysFunc)
	sysfunc.Id = id

	err := orm.NewOrm().Read(sysfunc)
	if err != nil {
		return nil, err
	}
	return sysfunc, nil
}
