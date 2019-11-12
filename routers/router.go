package routers

import (
	"beego_blogweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //注册
    beego.Router("/register", &controllers.RegisterController{})
}
