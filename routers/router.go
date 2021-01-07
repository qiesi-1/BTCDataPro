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
    //因为qq快速登录的返回code不能改动只能修改路由名，原本为index
    beego.Router("/qqLogin",&controllers.IndexController{})//
	//发送验证码接口
	beego.Router("/send_sms",&controllers.SendSmsController{})
    //短信登录接口
    beego.Router("/login_sms",&controllers.LoginSmsController{})
    //qq二维码扫码登录
    beego.Router("/qq_login",&controllers.QQLoginController{})
    //qq邮箱验证码登录
    beego.Router("/email_login",&controllers.EmailLoginController{})
    //发送邮箱验证码
    beego.Router("/email",&controllers.EmailLoginController{})
    //qq邮箱保存到数据库
    beego.Router("/email_login_second",&controllers.EmailLoginSecondController{})
    //找回密码
    beego.Router("/retrieve_pwd",&controllers.RetrievePwdController{})
    //主页面的所有页面展示接口
    beego.Router("/getblockbyhash.html",&controllers.IndexGetBlockDataController{})
	beego.Router("/getbestblockhash.html",&controllers.IndexGetBestBlockHashController{})
	beego.Router("/getblockchaininfo.html",&controllers.IndexGetBlockChaininfoController{})
	beego.Router("/getblockcount.html",&controllers.IndexGetBlockCountController{})
	beego.Router("/getblockhash.html",&controllers.IndexGetBlockHashController{})
	beego.Router("/getdifficulty.html",&controllers.IndexGetDifficultyController{})
	beego.Router("/getnewaddress.html",&controllers.IndexGetNewAddressController{})

}
