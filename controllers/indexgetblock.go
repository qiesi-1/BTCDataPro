package controllers

import (
	"BTCDataPro/tools"
	"github.com/astaxie/beego"
)

type IndexGetBlockDataController struct {
	beego.Controller
}

func (i *IndexGetBlockDataController) Get(){
	i.TplName="getblockbyhash.html"

	result,_:=tools.GetBlockByHash("000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f")

	i.Data["Hash"]=result.Hash
	i.Data["Height"]=result.Height
	i.Data["Difficulty"]=result.Difficulty
	i.Data["Size"]=result.Size
	i.Data["Nonce"]=result.Nonce
	i.Data["Ntx"]=result.Ntx

}
