package user

import (
	"github.com/astaxie/beego"
	"go_beego/models/user"
	"fmt"
)


type UserController struct {
	beego.Controller
}

func (this *UserController) Login() {
	name := this.GetString("name")
	password := this.GetString("password")

	if name == "" {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "name错误"}
		this.ServeJSON()
		return
	}

	if password == "" {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "password错误"}
		this.ServeJSON()
		return
	}

	user := userModel.User{}

	if code, err := user.FindByName(name); code != 0 || err != nil {
		this.Data["json"] = map[string]interface{}{"code": code, "message": err}
		this.ServeJSON()
		return
	}
	fmt.Println(user)

	if user.Password != password {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "password错误"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录成功"}
	this.ServeJSON()
}