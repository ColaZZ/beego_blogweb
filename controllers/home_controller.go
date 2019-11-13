package controllers

import (
	"beego_blogweb/models"
	"fmt"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	tag := c.GetString("tag")
	page, _ := c.GetInt("page")
	var articleList []models.Article

	if len(tag) > 0 {
		articleList, _ := models.QueryArticlesWithTag(tag)
		c.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}
		articleList, _ = models.FindArticleWithPage(page)
		c.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		c.Data["HasFooter"] = true
	}

	c.Data["Content"] = models.MakeHomeBlocks(articleList, c.IsLogin)

	fmt.Println("IsLogin:", c.IsLogin, c.Loginuser)
	c.TplName = "home.html"
}
