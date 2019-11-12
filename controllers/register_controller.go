package controllers

import (
	"fmt"
	"log"
	"time"

	"beego_blogweb/models"
	"beego_blogweb/utils"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	fmt.Println(username, password, repassword)

	//注册之前先判断用户是否已经被注册
	id := models.QueryUserWithUsername(username)
	log.Println("id", id)
	if id > 0 {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在"}
		c.ServeJSON()
		return
	}

	//注册用户名和密码
	password = utils.MD5(password)
	fmt.Println("MD5加密后", password)

	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InertUser(user)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	c.ServeJSON()
}
