package warn

import (
	"github.com/astaxie/beego"
)

type WarnController struct {
	beego.Controller
}

func (this *WarnController) Post() {
	var res string = "Warnning test!!!"
	this.Ctx.Output.Body([]byte(res))
}
