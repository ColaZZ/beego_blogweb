package controllers

import (
	"beego_blogweb/models"
	"fmt"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	var articleList []models.Article
	articleList, _ = models.FindArticleWithPage(page)
	c.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
	c.Data["HasFooter"] = true
	c.Data["Content"] = models.MakeHomeBlocks(articleList, c.IsLogin)

	fmt.Println("IsLogin:", c.IsLogin, c.Loginuser)
	c.TplName = "home.html"
}
