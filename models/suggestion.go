package models

import "github.com/astaxie/beego/orm"

type Suggestion struct {
	Content string `orm:"pk"`
	UserId  string `orm:"size(12)"`
	Email   string `orm:"size(30)"`
}

func AddSuggestion(uid, suggestion, email string) bool {
	o := orm.NewOrm()
	s := Suggestion{
		Content: suggestion,
		UserId:  uid,
		Email:   email,
	}
	_, err := o.Insert(&s)
	if err == nil {
		return true
	}
	return false
}
