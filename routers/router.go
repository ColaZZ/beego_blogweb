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

    //写文章
    beego.Router("/article/add", &controllers.AddArtileController{})
	//显示文章内容
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	//更新文章
	beego.Router("/article/update", &controllers.UpdateArticleController{})
    //删除文章
    beego.Router("/article/delete", &controllers.DeleteArticleController{})
}
