package controllers

import (
	"api/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type SuggestionController struct {
	beego.Controller
}

// @Title AddSuggestion
// @Description 添加建议
// @Param   Authorization  header    string  true		"token"
// @Param	requestBody	 body 	{}	 true		"email & suggestion"
// @Success 100 添加成功
// @Failure 101 添加失败
// @router / [post]
func (s *SuggestionController) AddSuggestion() {
	token, err := s.ParseToken()
	if err != "" {
		s.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		uid := claims["id"].(string)
		var data map[string]interface{}
		json.Unmarshal(s.Ctx.Input.RequestBody, &data)
		suggestion := data["suggestion"].(string)
		email := data["email"].(string)
		if !ok {
			s.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.AddSuggestion(uid, suggestion, email)
			s.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	s.ServeJSON()
}

//验证token
func (s *SuggestionController) ParseToken() (t *jwt.Token, err string) {
	authString := s.Ctx.Input.Header("Authorization")
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
