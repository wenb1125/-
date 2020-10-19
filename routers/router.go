package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //用户注册的接口请求
    beego.Router("/user_register",&controllers.RegisterController{})

    //直接登录的页面请求接口
    beego.Router("/login.html",&controllers.LoginController{})

    //用户登录请求接口
    beego.Router("/user_login",&controllers.LoginController{})

    //用户上传文件
	beego.Router("/upload", &controllers.UploadController{})

    //转跳新增页面
    beego.Router("/upload_file.html",&controllers.UploadController{})
}
