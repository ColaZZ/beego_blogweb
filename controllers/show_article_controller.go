package controllers

import (
	"beego_blogweb/utils"
	"gin_blogweb/models"
	"github.com/astaxie/beego"
	"strconv"
)

type ShowAritcleController struct {
	beego.Controller
}

func (c *ShowAritcleController) Get () {
	idStr := c.Ctx.Input.Param("id")
	id, _ := strconv.Atoi(idStr)

	//获取id所对应的的文章信息
	article := models.QueryArticleWithId(id)
	c.Data["Title"] = article.Title
	//c.Data["Content"] = article.Content
	c.Data["Content"] = utils.SwitchMarkdownToHtml(article.Content)
	c.TplName = "show_article.html"
}
