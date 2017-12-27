package models

import (
	"github.com/astaxie/beego/orm"
)

type SysUser struct {
	Id         int
	Username   string
	Nick       string
	Sex        int
	Password   string
	Mail       string
	Tel        string
	Roleid     int
	Status     int
	UpdateTime int64
	CreateTime int64
}

func (u *SysUser) TableName() string {
	return "sys_user"
}

func (u *SysUser) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func UserGetList(page, pageSize int, filters ...interface{}) ([]*SysUser, int64) {
	offset := (page - 1) * pageSize

	list := make([]*SysUser, 0)

	query := orm.NewOrm().QueryTable("sys_user")

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

func UserGetById(id int) (SysUser, error) {

	o := orm.NewOrm()
	user := SysUser{Id: id}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		return user, orm.ErrNoRows
	} else if err == orm.ErrMissPK {
		return user, orm.ErrMissPK
	} else {
		//fmt.Println(user.Id, user.Name)
	}

	return user, nil
}
