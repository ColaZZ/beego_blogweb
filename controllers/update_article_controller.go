package controllers

import (
	"beego_blogweb/models"
	"github.com/astaxie/beego"
)

type UpdateArticleController struct {
	beego.Controller
}

func (c *UpdateArticleController) Get() {
	id, _ := c.GetInt("id")
	//获取id对应的文章信息
	art := models.QueryArticleWithId(id)
	c.Data["Title"] = art.Title
	c.Data["Tags"] = art.Tags
	c.Data["Short"] = art.Short
	c.Data["Content"] = art.Content
	c.Data["Id"] = art.Id
	c.TplName = "write_article.html"
}

func (c *UpdateArticleController) Post() {
	id, _ := c.GetInt("id")
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")

	//实例化model，修改数据库
	art := models.Article{
		Id:         id,
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     "",
		CreateTime: 0,
	}
	_, err := models.UpdateArticle(art)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code":1, "message":"更新成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code":0, "message":"更新失败"}
	}
	c.ServeJSON()
}
