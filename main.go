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
}

