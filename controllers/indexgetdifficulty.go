package controllers

import "github.com/astaxie/beego"

type IndexGetDifficultyController struct {
	beego.Controller
}

func (i *IndexGetDifficultyController) Get(){
	i.TplName="getdifficulty.html"
}
