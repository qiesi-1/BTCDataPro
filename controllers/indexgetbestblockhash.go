package controllers

import "github.com/astaxie/beego"

type IndexGetBestBlockHashController struct {
	beego.Controller
}

func (i *IndexGetBestBlockHashController) Get(){
	i.TplName="getbestblockhash.html"
}
