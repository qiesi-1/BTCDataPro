package controllers

import "github.com/astaxie/beego"

type IndexGetBlockChaininfoController struct {
	beego.Controller
}

func (i *IndexGetBlockChaininfoController) Get(){
	i.TplName="getblockchaininfo.html"
}
