package controllers

import (
	"beego_blogweb/models"
	"github.com/astaxie/beego/logs"
)

type AlbumController struct {
	BaseController
}

func (c *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		logs.Error(err)
	}
	c.Data["Album"] = albums
	c.TplName ="album.html"
}
