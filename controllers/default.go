package controllers

import (
	"beegonew0604/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
	//第二种方法name:=c.GetString("name")
	//age1,err:=c.GetInt("age")

	//JSON格式提交:{
	//    "name" : "tengyuan",
	//    "age" : 18,
	//    "sex" ："famale"
	//}
	//后台处理

	//获取get类型请求的请求参数
	name := c.Ctx.Input.Query("name")
	fmt.Println("名字为:",name)
	age := c.Ctx.Input.Query("age")
	fmt.Println("年龄是:",age)
	sex := c.Ctx.Input.Query("sex")
	fmt.Println("性别是:",sex)

	if name != "藤原千花" || age != "18" {
		c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("数据验证成功"))
	//c.ctx.writeString()直接写入字符串

	c.Data["Website"] = "wwww.baidu.com"
	c.Data["Email"] = "hanhan@qq.com"
	c.TplName = "index.tpl"
}

//该post方法是处理post的请求的时候要调用的方法
func (c *MainController) Post(){
	//request 请求,response响应
	fmt.Println("post请求")
	/*user := c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为:",user)
	psd := c.Ctx.Request.FormValue("psd")
	fmt.Println("密码是:",psd)
	//与固定值比较
	if user != "藤原千花"  ||  psd != "123456"{
		//失败页面
		c.Ctx.ResponseWriter.Write([]byte("对不起数据不正确"))
	}else {
		c.Ctx.ResponseWriter.Write([]byte("恭喜你,数据正确"))
	}*/
	//后台处理
	body := c.Ctx.Request.Body
	dataByes,err:=ioutil.ReadAll(body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败")
		return
	}
	//json包解析
	var person models.Person
	json.Unmarshal(dataByes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败")
		return
	}
	fmt.Println("用户名",person.Name,"年龄",person.Age,"性别",person.Sex)
	c.Ctx.WriteString("用户名:"+person.Name)

}
