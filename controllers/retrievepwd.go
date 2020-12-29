package controllers

import "github.com/astaxie/beego"

type RetrievePwdController struct {
	beego.Controller
}


func (r *RetrievePwdController) Get(){

	r.TplName="retrieve_pwd.html"

}

