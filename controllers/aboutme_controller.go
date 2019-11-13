package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.Data["wechat"] = "微信："
	c.Data["qq"] = "QQ："
	c.Data["tel"] = "Tel："
	c.TplName = "aboultme.html"
}
