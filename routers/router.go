package routers

import (
	"beego_blogweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})

    //注册
    beego.Router("/register", &controllers.RegisterController{})

    //登录
    beego.Router("/login", &controllers.LoginContoller{})

    // 登出
    beego.Router("/exit", &controllers.ExitContoller{})
}
