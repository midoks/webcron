package models

import (
	"github.com/astaxie/beego/orm"
)

type SysLog struct {
	Id      int
	Uid     int
	Type    int
	Msg     string
	AddTime int64
}

func (u *SysLog) TableName() string {
	return "sys_logs"
}

func (u *SysLog) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func LogGetList(page, pageSize int, filters ...interface{}) ([]*SysLog, int64) {
	offset := (page - 1) * pageSize

	list := make([]*SysLog, 0)

	query := orm.NewOrm().QueryTable("sys_logs")

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			print(filters[k].(string), filters[k+1])
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}