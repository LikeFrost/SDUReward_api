package models

import (
	"os"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        string `orm:"column(SDU_number);pk;size(12)"`
	Password  string `orm:"size(255)"`
	Username  string `orm:"size(20)"`
	Telephone string `orm:"size(11)"`
}

//注册
func LogUp(id, password string) (code int, msg string) {
	o := orm.NewOrm()
	u := User{Id: id}
	err := o.Read(&u)
	if err == nil {
		return 101, "用户已存在,请直接登录!"
	} else {
		u.Password = password
		_, err = o.Insert(&u)
		if err == nil {
			return 100, "注册成功!"
		} else {
			return 102, "注册失败,请稍后再试!"
		}
	}
}

//登录
func Login(id, password string) (code int, msg string) {
	o := orm.NewOrm()
	u := User{Id: id}
	err := o.Read(&u)
	if err != nil {
		return 101, "用户不存在,请注册后登录!"
	} else if u.Password != password {
		return 102, "密码错误,请检查后重试!"
	}
	return 100, "登录成功!"
}

//更新用户信息
func UpdateUser(id, password, name, telephone string) (code int, msg string) {
	o := orm.NewOrm()
	u := User{Id: id}
	err := o.Read(&u)
	if err == nil {
		if password != "" {
			u.Password = password
		}
		if name != "" {
			u.Username = name
		}
		if telephone != "" {
			u.Telephone = telephone
		}
		_, err := o.Update(&u)
		if err == nil {
			return 100, "更新信息成功"
		} else {
			return 101, "更新信息失败"
		}
	}
	return 105, "用户不存在"
}

//获取用户信息
func GetUser(id string) (u User, code int, msg string) {
	o := orm.NewOrm()
	u = User{Id: id}
	err := o.Read(&u)
	if err == nil {
		return u, 100, "获取用户信息成功"
	}
	return u, 101, "获取用户信息失败"
}
func DeleteUser(id string) (code int, msg string) {
	o := orm.NewOrm()
	u := User{Id: id}
	r := []Reward{}
	err := o.Read(&u)
	if err == nil {
		_, er := o.QueryTable("reward").Filter("UserId__in", id).All(&r)
		if er == nil {
			for _, x := range r {
				os.Remove(x.Img)
				_, err = o.Delete(&x)
			}
			_, e := o.Delete(&u)
			if e == nil {
				return 100, "删除成功"
			}
		}
	}
	return 101, "删除失败"
}

func GetAllUsers(id string) (u []*User, code int, msg string) {
	o := orm.NewOrm()
	if id == "999999999" {
		_, err := o.QueryTable("user").Exclude("Id__in", "999999999").All(&u)
		if err == nil {
			return u, 100, "获取全部用户成功"
		} else {
			return u, 101, "获取全部用户失败"
		}
	} else {
		return u, 103, "权限不足"
	}

}
