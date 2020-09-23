package routers

import (
	"beegonew0604/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/login", &controllers.MainController{})
}
