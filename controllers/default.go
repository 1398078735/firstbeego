package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "wwww.baidu.com"
	c.Data["Email"] = "hanhan@qq.com"
	c.TplName = "index.tpl"
}
