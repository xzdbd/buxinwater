package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xzdbd/squeak-fuxinwater/models"
)

type MainController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	var loginStatus bool
	user := models.Userinfo{}
	if err := c.ParseForm(&user); err != nil {
		beego.Error(err.Error())
	}

	if user.Username == "admin" && user.Password == "admin" {
		loginStatus = true
	}

	if loginStatus {
		c.Redirect("/index", 302)
	} else {
		c.Data["isLoginFail"] = true
	}
	c.TplName = "login.tpl"
}
