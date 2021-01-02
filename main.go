package main

import (
	"BTCDataPro/db_mysql"
	_ "BTCDataPro/routers"
	"BTCDataPro/tools"
	"github.com/astaxie/beego"
)

func main() {
	tools.Getblockinfo()

	db_mysql.DbConnect()
	//静态资源文件路径
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()

}

