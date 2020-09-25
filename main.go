package main

import (
	_ "beegonew0604/routers"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//解析配置选项
	config :=beego.AppConfig //定义config变量,接收并赋值为全局配置变量
	//获取配置选项
	appName := config.String("appname")
	fmt.Println("项目应用名称",appName)
	port,err := config.Int("httpport")
	if err != nil {
		panic("项目配置信息解析错误")
	}
	fmt.Println("应用的监听端口",port)

	dbdriver := config.String("db_driver")
	dbuser := config.String("db_user")
	dbpassword:= config.String("db_password")
	dbip := config.String("db_ip")
	dbname := config.String("db_name")

	db,err:=sql.Open(dbdriver,dbuser+":"+dbpassword+"@tcp("+dbip+")/"+dbname+"?charset=utf8")
	if err != nil {
		panic("数据连接打开失败")
	}
	fmt.Println(db)

	beego.Run()
	//beego框架的配置文件操作:
	//全局应用配置结构体变量:beego.appconfig
}

