package controllers

import (
	"BTCDataPro/tools"
	"github.com/astaxie/beego"
)

type IndexGetBlockChaininfoController struct {
	beego.Controller
}

func (i *IndexGetBlockChaininfoController) Get(){
	i.TplName="getblockchaininfo.html"

	result,_:=tools.GetBlockChainInfo()

	i.Data["Chain"]=result.Chain
	i.Data["Hash"]=result.Bestblockhash
	i.Data["Difficulty"]=result.Difficulty
	i.Data["Size"]=result.Size_on_disk
	i.Data["Count"]=result.Blocks
}
