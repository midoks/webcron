package models

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type AppCronLog struct {
	Id          int
	CronId      int
	Output      string
	Error       string
	ProcessTime int
	Status      int
	CreateTime  int64
}

func (u *AppCronLog) TableName() string {
	return TableName("cron_log")
}

func (u *AppCronLog) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func CronLogGetList(page, pageSize int, filters ...interface{}) ([]*AppCronLog, int64) {
	offset := (page - 1) * pageSize

	list := make([]*AppCronLog, 0)

	query := orm.NewOrm().QueryTable(TableName("cron_log"))

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

func CronLogGetById(id int) (*AppCronLog, error) {

	u := new(AppCronLog)
	err := orm.NewOrm().QueryTable(TableName("cron_log")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func CronLogGetByName(name string) (*AppCronLog, error) {

	u := new(AppCronLog)
	err := orm.NewOrm().QueryTable(TableName("cron_log")).Filter("name", name).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func CronLogAdd(cron *AppCronLog) (int64, error) {
	if cron.CreateTime == 0 {
		cron.CreateTime = time.Now().Unix()
	}
	return orm.NewOrm().Insert(cron)
}

func CronLogDelById(id int) (int64, error) {
	return orm.NewOrm().Delete(&AppCronLog{Id: id})
}
