package controllers

import (
	"beego_blogweb/models"
	"github.com/astaxie/beego"
	"log"
)

type DeleteArticleController struct {
	beego.Controller
}

func (c *DeleteArticleController) Get() {
	artID, err := c.GetInt("artId")
	if err != nil {
		log.Println(err)
	}
	_, err = models.DeleteArticle(artID)
	if err != nil {
		log.Println(err)
	}
	c.Redirect("/", 302)
}
