package controllers

import (
	"BTCDataPro/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/skip2/go-qrcode"
	"math/rand"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get(){
	l.TplName = "login.html"
}

func (l *LoginController) Post(){
	//二维码登录
	rand.Seed(time.Now().UnixNano())
	time := rand.Intn(1000)
	err := qrcode.WriteFile(" 扫码呀憨憨 "+string(time), qrcode.Medium, 256, "./static/img/qrcode.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//解析客户端用户提交的登录数据
	var user models.Users
	//usename := l.Ctx.Request.PostForm.Get("user")
	//pwd := l.Ctx.Request.PostForm.Get("pwd")
	//fmt.Println(usename,pwd)
	//user.UserName = usename
	//user.Pwd = pwd
	err = l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("抱歉,用户登录信息解析错误")
		return
	}
	//根据解析到的数据执行数据库查询操作
	_,err=user.QueryUsers()
	//判断查询数据库的结果
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉,用户登录错误")
		return
	}

	l.TplName="index.html"//跳转到主页面
}
