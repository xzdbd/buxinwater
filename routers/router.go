package routers

import (
	"github.com/astaxie/beego"
	"github.com/xzdbd/squeak-fuxinwater/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/index", &controllers.MainController{})
}
