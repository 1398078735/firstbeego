package db_mysql

import (
	"beegonew0604/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
)

var Db *sql.DB
func Connect(){
	config :=beego.AppConfig //定义config变量,接收并赋值为全局配置变量
	//获取配置选项
	appName := config.String("appname")
	fmt.Println("项目应用名称",appName)
	port,err := config.Int("httpport")
	if err != nil {
		panic("项目配置信息解析错误")
	}
	fmt.Println("应用的监听端口",port)

	driver := config.String("db_driver")
	dbuser := config.String("db_user")
	dbpassword:= config.String("db_password")
	dbip := config.String("db_ip")
	dbname := config.String("db_name")

	//链接数据库 函数调用sql.open 该函数有两个参数
	//数据库驱动 mysql,sql server...
	//要求:数据库创建时,使用utf8编码作为默认的编码

	db,err := sql.Open(driver,dbuser+":"+dbpassword+"@tcp("+dbip+")/"+dbname+"?charset=utf8")
	if err != nil {
		//早发现,早解决
		panic("数据连接失败")//让程序进入恐慌状态
	}
	Db = db
	fmt.Println(db)

	//代码封装:可以将重复的代码或者功能相对比较独立的代码，进行封装，以
	//函数的形式进行封装，变成一个代码块或者是功能包，供使用者进行调用

}

//将用户信息保存到数据库中的函数

func AddUser(u models.User) (int64,error){
	//将密码进行hash计算，得到密码的hash值,然后再存
	md5Hash:=md5.New()
	md5Hash.Write([]byte(u.Nick))
	NickBytes:=md5Hash.Sum(nil)
	u.Nick = hex.EncodeToString(NickBytes)

	result,err:=Db.Exec("insert into user(Name,Birthday,Address,Nick)" +
	"values(?,?,?,?)",u.Name,u.Birthday,u.Address,u.Nick)
	if err != nil {
		return -1,err
	}
	row,err:=result.RowsAffected()//影响了的几行
	if err != nil {
		return -1,err
	}
	return row,err
}