package routers

import (
	"go_beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/warnning", &controllers.WarnningController{})
}
