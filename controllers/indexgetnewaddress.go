package controllers

import (
	"BTCDataPro/tools"
	"github.com/astaxie/beego"
)

type IndexGetNewAddressController struct {
	beego.Controller
}

func (i *IndexGetNewAddressController) Get(){
	i.TplName= "getnewaddress.html"

	result,_:=tools.GetNewAddress()

	i.Data["Address"]=result
}
