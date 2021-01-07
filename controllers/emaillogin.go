package controllers

import (
	"BTCDataPro/models"
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
)

type EmailLoginController struct {
	beego.Controller
}

func (e *EmailLoginController) Get(){
	e.TplName="loginemail.html"
}

//发送邮箱给用户
//用到goail库，用于发送邮件是个很强的库
func (e *EmailLoginController) Post(){
	var email models.EmailUser
	err := e.ParseForm(&email)
	if err != nil {
		e.Ctx.WriteString("抱歉,邮箱信息解析错误")
		return
	}
	//发送邮箱
	// 初始化
	m := gomail.NewMessage()
	// 发邮件的地址
	m.SetHeader("From", "tengyuanqianhua5@qq.com")
	// 给谁发送，支持多个账号
	m.SetHeader("To", email.EmailName)
	// 邮件标题
	m.SetHeader("Subject", "你就是区块链未来的大佬嘛！欢迎你来到我们的系统！")
	// 邮件正文，支持 html
	m.SetBody("text/html", "欢迎你注册成为我们区块链节点通信的会员，验证码是:比特币 人类文明已经从“身份社会”进化到了“契约社会”，而区块链有望带领人类从“契约社会”过渡到“智能合约”的社会。" +
		"真正的验证码是:你的生日")
	// stmp服务，端口号，发送邮件账号，发送账号密码
	d := gomail.NewDialer("smtp.qq.com", 465, "tengyuanqianhua5@qq.com", "psvmkudaqddthcdg")
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	gomail.NewMessage(gomail.SetEncoding(gomail.Base64))

	e.Data["Email"]=email.EmailName
	e.TplName="loginemail_second.html"
}