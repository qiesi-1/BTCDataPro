package controllers

import (
	"BTCDataPro/tools"
	"github.com/astaxie/beego"
)

type IndexGetBestBlockHashController struct {
	beego.Controller
}

func (i *IndexGetBestBlockHashController) Get(){
	i.TplName="getbestblockhash.html"

	result,_:=tools.GetBestBlockHash()

	i.Data["Hash"]=result
}
