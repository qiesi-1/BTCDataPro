package controllers

import (
	"BTCDataPro/models"
	"fmt"
	"github.com/astaxie/beego"
)

type EmailLoginSecondController struct {
	beego.Controller
}

//用于将邮箱保存到数据库
func (e *EmailLoginSecondController) Post(){
	var email models.EmailUser
	err := e.ParseForm(&email)
	if err != nil {
		e.Ctx.WriteString("抱歉,邮箱信息解析错误")
		return
	}
	//将邮箱信息保存到数据库
	_,err = email.AddEmailUser()
	if err != nil {
		fmt.Println(err.Error())
		e.Ctx.WriteString("邮箱登录信息错误")
		return
	}

	e.Data["Email"] = email.EmailName
	e.TplName="index.html"
}