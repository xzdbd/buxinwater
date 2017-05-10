package routers

import (
	"github.com/astaxie/beego"
	"github.com/xzdbd/squeak-fuxinwater/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
