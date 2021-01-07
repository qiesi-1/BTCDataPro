package controllers

import (
	"BTCDataPro/tools"
	"github.com/astaxie/beego"
)

type IndexGetBlockCountController struct {
	beego.Controller
}

func (i *IndexGetBlockCountController) Get(){
	i.TplName="getblockcount.html"

	result,_:=tools.GetBlockCount()

	i.Data["Count"]=result
}
