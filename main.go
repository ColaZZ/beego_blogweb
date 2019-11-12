package main

import (
	_ "beego_blogweb/routers"
	"beego_blogweb/utils"
	"github.com/astaxie/beego"
)

func main() {
	utils.InitMysql()
	beego.Run()

}

