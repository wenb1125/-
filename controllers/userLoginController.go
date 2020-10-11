package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	//设置login.html为模板文件
	l.TplName = "login.html"
}
//用户登录接口
func (l *LoginController) Post() {
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("抱歉，用户信息解析失败，请重试！")
		return
	}
	//查询数据库用户信息
	_, err = user.QueryUser()
	if err != nil {
		l.Ctx.WriteString("抱歉，用户登录失败，请重试！")
		return
	}
	//登录成功
	l.Data["phone"] = user.Phone
	l.TplName = "home.html"
}