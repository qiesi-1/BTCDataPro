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
    //主页面接口
    beego.Router("/index",&controllers.IndexController{})
	//发送验证码注册接口
	beego.Router("/send_sms",&controllers.SendSmsController{})
    //手机短信登录接口
    beego.Router("/login_sms",&controllers.LoginSmsController{})
}
