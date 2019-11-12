package controllers

import (
	"beego_blogweb/models"
	"beego_blogweb/utils"
	"github.com/astaxie/beego"
)

type LoginContoller struct {
	beego.Controller
}

func (c *LoginContoller) Get () {
	c.TplName = "login.html"
}

func (c *LoginContoller) Post () {
	username := c.GetString("username")
	password := c.GetString("password")

	password = utils.MD5(password)
	id := models.QueryUserWithParam(username, password)
	if id > 0 {
		c.Data["json"] = map[string]interface{}{"code":1, "message":"登录成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code":0, "messsage":"登录失败"}
	}
	c.ServeJSON()
}
