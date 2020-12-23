package controllers

import (
	"BTCDataPro/models"
	"fmt"
	"github.com/astaxie/beego"
)

type ReisterController struct {
	beego.Controller
}

func (r *ReisterController) Get(){
	//解析用户登录请求数据
	var user models.Users
	err := r.ParseForm(&user)
	if user.Pwd != user.Rpwd {
		r.Ctx.WriteString("输入的两次密码不一致请重试")
		return
	}
	if err != nil {
		r.Ctx.WriteString("抱歉,数据错误")
		return
	}

	_,err = user.AddUsers()
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("注册用户信息错误")
		return
	}

	r.TplName = "login.html"
}
