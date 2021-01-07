package controllers

import (
	"BTCDataPro/tools"
	"github.com/astaxie/beego"
)

type IndexGetBlockHashController struct {
	beego.Controller
}

func (i *IndexGetBlockHashController) Get(){
	i.TplName="getblockhash.html"

	result,_:=tools.GetBlockHashByHeight(0)

	i.Data["Hash"]=result
}
