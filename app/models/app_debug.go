package models

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type AppDebug struct {
	Id         int
	Name       string
	Desc       string
	Type       int
	ServerId   string
	Status     int
	UpdateTime int64
	CreateTime int64
}

func (u *AppDebug) TableName() string {
	return TableName("debug")
}

func (u *AppDebug) Update(fields ...string) error {
	u.UpdateTime = time.Now().Unix()
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func DebugGetList(page, pageSize int, filters ...interface{}) ([]*AppDebug, int64) {
	offset := (page - 1) * pageSize

	list := make([]*AppDebug, 0)

	query := orm.NewOrm().QueryTable(TableName("item"))

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			// print(filters[k].(string), filters[k+1])
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}

func DebugGetById(id int) (*AppDebug, error) {

	u := new(AppDebug)
	err := orm.NewOrm().QueryTable(TableName("item")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func DebugGetByName(name string) (*AppDebug, error) {

	u := new(AppDebug)
	err := orm.NewOrm().QueryTable(TableName("item")).Filter("name", name).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func DebugDelById(id int) (int64, error) {
	return orm.NewOrm().Delete(&AppDebug{Id: id})
}
