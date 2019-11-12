package controllers

type ExitContoller struct {
	BaseController
}

func (c *ExitContoller) Get() {
	c.DelSession("loginuser")
	c.Redirect("/", 302)
}
