package main

import (
	"beegonew0604/db_mysql"
	_ "beegonew0604/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	//1，链接数据库
	db_mysql.Connect()

	beego.Run()
	//beego框架的配置文件操作:
	//全局应用配置结构体变量:beego.appconfig
	//c.ctx.writestring(string):向客户端返回字符串
	//c.ctx.response.write:向客户端返回字节切片字符串
	//为了返回更丰富的数据类型，在beego中，提供了一种数据返回格式:
	//Json数据格式
	//后台返回json格式:
	//json{}单个结构体,【】代表多个同类型的结构体，数组/切片
}

