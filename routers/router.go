package routers

import (
	"github.com/astaxie/beego"
	"go_beego/controllers/users"
	"go_beego/controllers/index"
	"go_beego/controllers/warn"
)

func init() {
	beego.Router("/", &index.IndexController{})
	beego.Router("/warnning", &warn.WarnController{})
	beego.Router("/user/login", &users.UserLoginController{})
}
