package controllers

import (
	"github.com/astaxie/beego"
)

type TuController struct {
	beego.Controller
}


//***********
//
// 执行任务
//
//**********
func (c *TuController) Index() {

    pdfPath := "正在处理任务..."

    
	c.Data["json"] = pdfPath
	c.ServeJSON()
}

