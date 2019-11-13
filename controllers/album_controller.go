package controllers

type AlbumController struct {
	BaseController
}

func (c *AlbumController) Get() {
	c.TplName ="album.html"
}
