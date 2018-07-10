package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type WarnningController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *WarnningController) Post() {
	// c.Ctx.WriteString("warnning test")
	c.Ctx.Output.Body([]byte("warnning !!!"))
}
