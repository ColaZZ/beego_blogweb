package controllers

import "fmt"

type HomeController struct {
	BaseController
}

func (c *HomeController) Get () {
	fmt.Println("IsLogin:", c.IsLogin, c.Loginuser)
	c.TplName = "home.html"
}
