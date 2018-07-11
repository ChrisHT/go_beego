package users

import (
	"github.com/astaxie/beego"
	"go_beego/services/users"
	"fmt"
)

type UserLoginController struct {
	beego.Controller
}

func (this *UserLoginController) Post() {
	name := this.GetString("name")
	password := this.GetString("password")

	if password == "" {
		this.Ctx.WriteString("password is required")
		return
	}
	fmt.Println(name, " ", password)
	res := userService.UserLoginService(name)
	this.Ctx.Output.Body([]byte(res))
}
