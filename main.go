package main

import (
	"BTCDataPro/db_mysql"
	_ "BTCDataPro/routers"
	"github.com/astaxie/beego"
)

func main() {

	db_mysql.DbConnect()
	//静态资源文件路径
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()
}

