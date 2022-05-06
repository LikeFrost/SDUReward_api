package controllers

import (
	"api/models"
	"crypto/md5"
	"encoding/hex"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type RewardController struct {
	beego.Controller
}

// @Title AddReward
// @Description 添加奖励记录
// @Param   Authorization  	header    	string  true		"token"
// @Param	requestBody	 body 	{}	 true		"tagId & type & name & grade & prize & score & time & img"
// @Success 100 添加奖励成功
// @Failure 101 添加奖励失败
// @Failure 102 token失效
// @Failure 103 图片上传失败
// @router / [post]
func (r *RewardController) Add() {
	token, err := r.ParseToken()
	if err != "" {
		r.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			r.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			uid := claims["id"].(string)
			tag := r.GetString("tag")
			rewardType := r.GetString("type")
			name := r.GetString("name")
			grade := r.GetString("grade")
			prize := r.GetString("prize")
			score, _ := r.GetInt("score")
			time := r.GetString("time")
			md5Data := md5.New()
			md5Data.Write([]byte(uid + time + tag + name))
			path := "img/" + hex.EncodeToString(md5Data.Sum(nil)) + ".jpg"
			e := r.SaveToFile("img", path)
			if e != nil {
				r.Data["json"] = map[string]interface{}{
					"code": 103,
					"msg":  "图片上传失败,请稍后再试",
				}
			} else {
				code, msg := models.AddReward(uid, name, time, path, tag, rewardType, grade, prize, score)
				r.Data["json"] = map[string]interface{}{
					"code": code,
					"msg":  msg,
				}
			}
		}
	}
	r.ServeJSON()
}

// @Title GetReward
// @Description 获取奖励记录
// @Param   Authorization  	header		string  	true		"token"
// @Param   rewardId        path        int  		true		"rewardId"
// @Success 100 获取成功
// @Failure 101 获取失败
// @Failure 102 token失效
// @router /:rewardId [get]
func (r *RewardController) GetReward() {
	token, err := r.ParseToken()
	if err != "" {
		r.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		rewardId, _ := r.GetInt(":rewardId")
		if !ok {
			r.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			reward, code, msg := models.GetReward(rewardId)
			if code == 100 && (reward.UserId == id || id == "999999999") {
				r.Data["json"] = map[string]interface{}{
					"code":   code,
					"msg":    msg,
					"reward": reward,
				}
			} else {
				r.Data["json"] = map[string]interface{}{
					"code": code,
					"msg":  msg,
				}
			}
		}
	}
	r.ServeJSON()
}

// @Title GetRewardByTag
// @Description 获取分类奖励记录
// @Param   Authorization  	header		string  	true		"token"
// @Param   tag        		path        string  	true		"tag"
// @Success 100 获取成功
// @Failure 101 获取失败
// @router /byTag/:tag [get]
func (r *RewardController) GetRewardByTag() {
	token, err := r.ParseToken()
	if err != "" {
		r.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		uid := claims["id"].(string)
		tag := r.GetString(":tag")
		if !ok {
			r.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			reward, code, msg := models.GetRewardByTag(tag, uid)
			r.Data["json"] = map[string]interface{}{
				"code":   code,
				"msg":    msg,
				"reward": reward,
			}
		}
	}
	r.ServeJSON()
}

// @Title GetAllReward
// @Description 获取所有奖励记录
// @Param   Authorization  	header		string  	true		"token"
// @Success 100 获取成功
// @Failure 101 获取失败
// @router / [get]
func (r *RewardController) GetAllReward() {
	token, err := r.ParseToken()
	if err != "" {
		r.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		if !ok {
			r.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			reward, code, msg := models.GetAllReward(id)
			r.Data["json"] = map[string]interface{}{
				"code":   code,
				"msg":    msg,
				"reward": reward,
			}
		}
	}
	r.ServeJSON()
}

// @Title GetAllReward
// @Description 获取某一用户奖励记录
// @Param   Authorization  	header		string  	true		"token"
// @Success 100 获取成功
// @Failure 101 获取失败
// @router /byUserId/:uid [get]
func (r *RewardController) GetUserReward() {
	token, err := r.ParseToken()
	if err != "" {
		r.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		if !ok {
			r.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else if id == "999999999" {
			uid := r.GetString(":uid")
			reward, code, msg := models.GetAllReward(uid)
			r.Data["json"] = map[string]interface{}{
				"code":   code,
				"msg":    msg,
				"reward": reward,
			}
		} else {
			r.Data["json"] = map[string]interface{}{
				"code": 103,
				"msg":  "权限不足",
			}
		}
	}
	r.ServeJSON()
}

// @Title DeleteReward
// @Description 删除奖励记录
// @Param   Authorization  	header		string  	true		"token"
// @Param   rewardId        path        int  		true		"rewardId"
// @Success 100 删除成功
// @Failure 101 删除失败
// @router /:rewardId [Delete]
func (r *RewardController) DeleteReward() {
	token, err := r.ParseToken()
	if err != "" {
		r.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		rewardId, _ := r.GetInt(":rewardId")
		if !ok {
			r.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.DeleteReward(rewardId, id)
			r.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	r.ServeJSON()
}

//验证token
func (r *RewardController) ParseToken() (t *jwt.Token, err string) {
	authString := r.Ctx.Input.Header("Authorization")
	if authString == "" {
		return t, "token失效"
	}
	token, e := jwt.Parse(authString, func(token *jwt.Token) (interface{}, error) {
		return []byte("zdgzsl"), nil
	})
	if e != nil {
		return token, "token失效"
	}
	if !token.Valid {
		return token, "token失效"
	}
	return token, ""
}
