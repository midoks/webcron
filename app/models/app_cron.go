package models

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type AppCron struct {
	Id         int
	Name       string
	Desc       string
	ItemId     int
	CronSpec   string
	Cmd        string
	Concurrent int
	ExecNum    int
	PrevTime   int
	Notify     int
	Timeout    int
	Status     int
	UpdateTime int64
	CreateTime int64
}

func (u *AppCron) TableName() string {
	return TableName("cron")
}

func (u *AppCron) Update(fields ...string) error {
	u.UpdateTime = time.Now().Unix()
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func CronGetList(page, pageSize int, filters ...interface{}) ([]*AppCron, int64) {
	offset := (page - 1) * pageSize

	list := make([]*AppCron, 0)

	query := orm.NewOrm().QueryTable(TableName("cron"))

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}

func CronGetById(id int) (*AppCron, error) {

	u := new(AppCron)
	err := orm.NewOrm().QueryTable(TableName("cron")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func CronGetByName(name string) (*AppCron, error) {

	u := new(AppCron)
	err := orm.NewOrm().QueryTable(TableName("cron")).Filter("name", name).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func CronDelById(id int) (int64, error) {
	return orm.NewOrm().Delete(&AppCron{Id: id})
}
