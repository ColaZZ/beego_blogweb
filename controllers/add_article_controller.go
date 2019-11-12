package controllers

import (
	"beego_blogweb/models"
	"github.com/astaxie/beego"
	"time"
)

type AddArtileController struct {
	beego.Controller
}

func (c *AddArtileController) Get() {
	c.TplName = "write_article.html"
}

func (c *AddArtileController) Post() {
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")
	author := c.GetString("author")

	article := models.Article{
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     author,
		CreateTime: time.Now().Unix(),
	}
	_, err := models.AddArticle(article)

	var response map[string]interface{}
	if err != nil {
		response = map[string]interface{}{"code":0, "message":"error"}
	} else {
		response = map[string]interface{}{"code":1, "message":"ok"}
	}
	c.Data["json"] = response
	c.ServeJSON()
}
