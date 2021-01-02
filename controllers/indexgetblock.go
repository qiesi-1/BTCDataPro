package controllers

import "github.com/astaxie/beego"

type IndexGetBlockController struct {
	beego.Controller
}

func (i *IndexGetBlockController) Get(){
	i.TplName="getblock.html"
}
