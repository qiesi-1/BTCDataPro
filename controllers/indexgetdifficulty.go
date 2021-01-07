package controllers

import (
	"BTCDataPro/tools"
	"github.com/astaxie/beego"
)

type IndexGetDifficultyController struct {
	beego.Controller
}

func (i *IndexGetDifficultyController) Get(){
	i.TplName="getdifficulty.html"

	result,_:=tools.GetDifficult()

	i.Data["Difficultly"]=result
}
