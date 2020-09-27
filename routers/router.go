package routers

import (
"beegonew0604/controllers"
"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register",&controllers.RegisterController{})

	beego.Router("/", &controllers.MainController{})
}
