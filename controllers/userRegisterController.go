package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	//1.解析数据请求、
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		//返回错误信息给浏览器，提示用户
		r.Ctx.WriteString("抱歉，解析数据错误，请重试！")
		return
	}
	//2.保存用户信息到数据库
	id, err := user.SaveUser()

	//3.返回前端结果(成功跳转页面，失败弹出错误信息)
	if err != nil {
		r.Ctx.WriteString("抱歉,用户注册失败,请重试！")
		return
	}
	fmt.Printf("影响了数据库%d行\n",id)
	//用户注册成功
	r.TplName = "login.html"
}