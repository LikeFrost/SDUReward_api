package models

import (
	"os"

	"github.com/astaxie/beego/orm"
)

type Reward struct {
	Id     int    `orm:"column(id);pk;auto;size(12)"`
	Tag    string `orm:"size(255)"`
	Type   string `orm:"size(255)"`
	Name   string `orm:"size(255)"`
	Grade  string `orm:"size(255)"`
	Prize  string `orm:"size(255)"`
	Score  int    `orm:"size(3)"`
	UserId string `orm:"size(12)"`
	Time   string `orm:"size(10)"`
	Img    string `orm:"size(255)"`
}

func AddReward(uid, name, time, path, tag, rewardType, grade, prize string, score int) (code int, msg string) {
	o := orm.NewOrm()
	r := Reward{
		Type:   rewardType,
		Tag:    tag,
		Name:   name,
		Grade:  grade,
		Prize:  prize,
		Score:  score,
		UserId: uid,
		Time:   time,
		Img:    path,
	}
	_, err := o.Insert(&r)
	if err != nil {
		return 101, "添加奖励失败"
	}
	return 100, "添加奖励成功"
}
func GetReward(id int) (r Reward, code int, msg string) {
	o := orm.NewOrm()
	r = Reward{Id: id}
	err := o.Read(&r)
	if err == nil {
		return r, 100, "获取奖励成功"
	}
	return r, 101, "获取奖励失败"
}
func GetRewardByTag(tag, uid string) (r []*Reward, code int, msg string) {
	o := orm.NewOrm()
	_, err := o.QueryTable("reward").Filter("UserId__in", uid).Filter("Tag__in", tag).All(&r)
	if err != nil {
		return r, 101, "获取分类奖励失败"
	} else {
		return r, 100, "获取分类奖励成功"
	}
}
func GetAllReward(uid string) (r []*Reward, code int, msg string) {
	o := orm.NewOrm()
	_, err := o.QueryTable("reward").Filter("UserId__in", uid).All(&r)
	if err != nil {
		return r, 101, "获取全部奖励失败"
	} else {
		return r, 100, "获取全部奖励成功"
	}
}

func DeleteReward(id int, uid string) (code int, msg string) {
	o := orm.NewOrm()
	r := Reward{Id: id}
	err := o.Read(&r)
	if err == nil && r.UserId == uid {
		os.Remove(r.Img)
		_, e := o.Delete(&r)
		if e == nil {
			return 100, "删除成功"
		}
	}
	return 101, "删除失败"
}
