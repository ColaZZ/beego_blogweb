package controllers

import (
	"beego_blogweb/models"
	"github.com/astaxie/beego"
)

type TagsController struct {
	beego.Controller
}

func (c *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")
	c.Data["Tags"] = models.HandlerTagListData(tags)
	c.TplName = "tags.html"
}
