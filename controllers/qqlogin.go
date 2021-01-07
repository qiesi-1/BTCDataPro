package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"net/url"
)

type QQLoginController struct {
	beego.Controller
}

const (
	AppId       = "101827468"
	redirectURI = "http://127.0.0.1:9090/qqLogin"
)

//获取认证的code
func (q *QQLoginController) Get () {
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", AppId)
	params.Add("state", "test")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), redirectURI)
	//登录成功，跳转到主页面
	loginURL := fmt.Sprintf("%s?%s", "https://graph.qq.com/oauth2.0/authorize", str)
	http.Redirect(q.Ctx.ResponseWriter,q.Ctx.Request,loginURL, http.StatusFound)
}

