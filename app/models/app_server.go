package models

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type AppServer struct {
	Id         int
	Ip         string
	Port       int
	Desc       string
	Type       int
	User       string
	Pwd        string
	PubKey     string
	Status     int
	UpdateTime int64
	CreateTime int64
}

func (u *AppServer) TableName() string {
	return TableName("server")
}

func (u *AppServer) Update(fields ...string) error {
	u.UpdateTime = time.Now().Unix()
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func ServerGetList(page, pageSize int, filters ...interface{}) ([]*AppServer, int64) {
	offset := (page - 1) * pageSize

	list := make([]*AppServer, 0)

	query := orm.NewOrm().QueryTable(TableName("server"))

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

func ServerGetById(id int) (*AppServer, error) {

	u := new(AppServer)
	err := orm.NewOrm().QueryTable(TableName("server")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func ServerGetByName(name string) (*AppServer, error) {

	u := new(AppServer)
	err := orm.NewOrm().QueryTable(TableName("server")).Filter("name", name).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func ServerDelById(id int) (int64, error) {
	return orm.NewOrm().Delete(&AppServer{Id: id})
}
