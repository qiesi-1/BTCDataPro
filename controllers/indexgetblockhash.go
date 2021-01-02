package controllers

import "github.com/astaxie/beego"

type IndexGetBlockHashController struct {
	beego.Controller
}

func (i *IndexGetBlockHashController) Get(){
	i.TplName="getblockhash.html"
}
