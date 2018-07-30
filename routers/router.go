package routers

import (
	"github.com/astaxie/beego"
	"go_beego/controllers/user"
	"go_beego/controllers/index"
)

func init() {
	beego.Router("/", &index.IndexController{}, "get:Index")
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/user",
			//beego.NSRouter("/register", &user.UserController{}, "post:Register"),
			beego.NSRouter("/login", &user.UserController{}, "post:Login"),
		),
	)
	beego.AddNamespace(ns)
}
