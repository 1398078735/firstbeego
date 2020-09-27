package controllers

import (
	"beegonew0604/db_mysql"
	"beegonew0604/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post(){
	//解析前端提交的用户注册的信息
	dataBytes,err:=ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("数据接收错误")
		return
	}
	var user models.User
	err = json.Unmarshal(dataBytes,&user)
	if err != nil {
		r.Ctx.WriteString("数据解析错误")
		return
	}
	//一切正常,将用户信息保存到数据库当中
	//直接调用保存数据的一个函数，并判断接收的结果
	row,err:=db_mysql.AddUser(user)
	if err != nil {
		r.Ctx.WriteString("注册用户失败")
		return
	}
	fmt.Println(row)
	r.Ctx.WriteString("恭喜注册用户信息成功")
}
