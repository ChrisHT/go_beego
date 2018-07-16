package users

import (
	"github.com/astaxie/beego"
	"go_beego/services/users"
	"fmt"
	"go_beego/models"
)

type UserLoginController struct {
	beego.Controller
}

func (this *UserLoginController) Post() {
	name := this.GetString("name")
	password := this.GetString("password")

	if password == "" {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "password错误"}
		this.ServeJSON()
		return
	}

	user := models.User{}
	code, err := user.FindByID(name)
	fmt.Println(err)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": code, "message": "can not find user"}
		this.ServeJSON()
		return
	}
	fmt.Println(name, " ", password)
	res := userService.UserLoginService(name)
	this.Data["json"] = map[string]interface{}{"code": 0, "message": res}
	this.ServeJSON()
}
