package routers

import (
	"BTCDataPro/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	//用户登录接口
	beego.Router("/login",&controllers.LoginController{})
    //用户注册接口
    beego.Router("/register",&controllers.ReisterController{})
}
