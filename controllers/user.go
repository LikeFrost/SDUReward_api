package controllers

import (
	"api/models"
	"encoding/json"
	"regexp"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Login
// @Description 登录
// @Param	requestBody	 body 	{}	 true		"id & password"
// @Success 100 登陆成功
// @Failure 101 用户不存在
// @Failure 102 密码错误
// @Failure 103 token生成失败
// @Failure 104 用户名格式错误
// @router /login [post]
func (u *UserController) Login() {
	var data map[string]interface{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &data)
	id := data["id"].(string)
	password := data["password"].(string)
	resultId, _ := regexp.MatchString(`^\d{9}$|^\d{12}$`, id)
	if !resultId {
		u.Data["json"] = map[string]interface{}{
			"code": 104,
			"msg":  "学号格式错误",
		}
	} else {
		code, msg := models.Login(id, password)
		if code == 100 {
			//创建token
			claims := make(jwt.MapClaims)
			claims["id"] = id
			claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte("zdgzsl"))
			if err == nil {
				u.Data["json"] = map[string]interface{}{
					"code":  code,
					"msg":   msg,
					"id":    id,
					"token": tokenString,
				}
			} else {
				u.Data["json"] = map[string]interface{}{
					"code":  103,
					"msg":   msg,
					"id":    id,
					"token": "生成失败",
				}
			}
		} else {
			u.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	u.ServeJSON()
}

// @Title LogUp
// @Description 注册
// @Param	requestBody	 body 	{}	 true		"id & password"
// @Success 100 注册成功
// @Failure 101 用户已经存在
// @Failure 102 注册失败，请稍后重试
// @Failure 103 token生成失败
// @Failure 104 学号格式错误
// @router /logup [post]
func (u *UserController) LogUp() {
	var data map[string]interface{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &data)
	id := data["id"].(string)
	password := data["password"].(string)
	resultId, _ := regexp.MatchString(`^\d{9}$|^\d{12}$`, id)
	if !resultId {
		u.Data["json"] = map[string]interface{}{
			"code": 104,
			"msg":  "学号格式错误",
		}
	} else {
		code, msg := models.LogUp(id, password)
		if code == 100 {
			//创建token
			claims := make(jwt.MapClaims)
			claims["id"] = id
			claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte("zdgzsl"))
			if err == nil {
				u.Data["json"] = map[string]interface{}{
					"code":  code,
					"msg":   msg,
					"id":    id,
					"token": tokenString,
				}
			} else {
				u.Data["json"] = map[string]interface{}{
					"code":  103,
					"msg":   msg,
					"id":    id,
					"token": "生成失败",
				}
			}
		} else {
			u.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	u.ServeJSON()
}

// @Title UpdateUser
// @Description 更新用户信息
// @Param   Authorization  header    string  true		"token"
// @Param	requestBody	 body 	{}	 true		"password & name & telephone"
// @Success 100 更新成功
// @Failure 101 更新失败
// @Failure 102 token失效
// @Failure 103 姓名格式错误
// @Failure 104 手机号格式错误
// @Failure 105 用户不存在
// @router / [post]
func (u *UserController) UpdateUser() {
	token, err := u.ParseToken()
	if err != "" {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		var data map[string]interface{}
		json.Unmarshal(u.Ctx.Input.RequestBody, &data)
		password := data["password"].(string)
		name := data["name"].(string)
		telephone := data["telephone"].(string)
		resultName, _ := regexp.MatchString(`^[\u4e00-\u9fa5]{2,10}$`, name)
		resultTelephone, _ := regexp.MatchString(`^1\d{10}$`, telephone)
		if !ok {
			u.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			if !resultName {
				u.Data["json"] = map[string]interface{}{
					"code": 103,
					"msg":  "姓名格式错误,仅可为2-10位汉字",
				}
			} else if !resultTelephone {
				u.Data["json"] = map[string]interface{}{
					"code": 104,
					"msg":  "手机号格式错误",
				}
			}
			code, msg := models.UpdateUser(id, password, name, telephone)
			u.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
				"Id":   id,
			}

		}
	}
	u.ServeJSON()
}

// @Title Get
// @Description 获取用户自身信息
// @Param   Authorization  header    string  true		"token"
// @Success 100 获取信息成功
// @Failure 101 获取信息失败
// @Failure 102 token失效
// @router / [get]
func (u *UserController) Get() {
	token, err := u.ParseToken()
	if err != "" {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
		u.ServeJSON()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	user, code, msg := models.GetUser(id)
	if !ok {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		u.Data["json"] = map[string]interface{}{
			"code": code,
			"msg":  msg,
			"user": user,
		}
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description 获取所有用户信息
// @Param   Authorization  header    string  true		"token"
// @Success 100 获取信息成功
// @Failure 101 获取信息失败
// @Failure 102 token失效
// @Failure 103 权限不足
// @router /all [get]
func (u *UserController) GetAllUsers() {
	token, err := u.ParseToken()
	if err != "" {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
		u.ServeJSON()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	user, code, msg := models.GetAllUsers(id)
	if !ok {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		u.Data["json"] = map[string]interface{}{
			"code": code,
			"msg":  msg,
			"user": user,
		}
	}
	u.ServeJSON()
}

// @Title DeleteUser
// @Description 删除用户
// @Param   Authorization  header    string  true		"token"
// @Success 100 删除成功
// @Failure 101 删除失败
// @Failure 102 token失效
// @Failure 103 权限不足
// @router /:uid [Delete]
func (u *UserController) DeleteUser() {
	token, err := u.ParseToken()
	if err != "" {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
		u.ServeJSON()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	uid := u.GetString(":uid")
	if !ok {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		if id == "999999999" {
			code, msg := models.DeleteUser(uid)
			u.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		} else {
			u.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "权限不足",
			}
		}
	}
	u.ServeJSON()
}

//验证token
func (u *UserController) ParseToken() (t *jwt.Token, err string) {
	authString := u.Ctx.Input.Header("Authorization")
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
