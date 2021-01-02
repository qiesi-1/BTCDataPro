package controllers

import "github.com/astaxie/beego"

type IndexGetBlockCountController struct {
	beego.Controller
}

func (i *IndexGetBlockCountController) Get(){
	i.TplName="getblockcount.html"
}
