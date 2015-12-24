package admin
import (
	"blog/controllers"
	"github.com/astaxie/beego"
	"blog/models"
	"blog/utils"
	"github.com/gogather/com"
	"fmt"
)
//注册
type RegisterController struct{
	controllers.BaseController
}
func (this *RegisterController)Get(){
	registerAble, err := beego.AppConfig.Bool("RegistorAble")
	if registerAble || err != nil{
		this.TplNames = "register.tpl"
	}else{
		this.Ctx.WriteString("register closed")
	}
}

func (this *RegisterController)Post(){
	registerAble, err := beego.AppConfig.Bool("RegistorAble")
	if err != nil {

	}else if !registerAble {
		this.Data["json"] = map[string]interface{}{"result":false, "msg":"register closed", "refer":"/"}
		this.ServeJson()
		return
	}
	username := this.GetString("username")
	password := this.GetString("password")

	if !utils.CheckUsername(username) {
		this.Data["json"] = map[string]interface{}{"result":false, "msg":"illegal username", "refer":"/"}
		this.ServeJson()
		return
	}
	id, err := models.AddUser(username, password)
	if err == nil && id > 0{
		this.Data["json"] = map[string]interface{}{"result":true, "msg":"register success", "refer":"/"}
	}else{
		this.Data["json"] = map[string]interface{}{"result": false, "msg":"register failed", "refer":"/"}
	}
	this.ServeJson()
}

//登陆
type LoginController struct {
	controllers.BaseController
}
func (this *LoginController)Get(){
	user := this.GetSession("username")
	fmt.Println(user)
	if user != nil {
		this.Redirect("/admin", 302)
	}else {
		this.TplNames = "login.tpl"
	}
}
func (this *LoginController)Post(){
	username := this.GetString("username")
	password := this.GetString("password")
	if username == "" || password == "" {
		this.Data["json"] = map[string]interface{}{"result":false, "msg":"invalid request", "refer":"/"}
		this.ServeJson()
		return
	}
	user, err := models.FindUser(username)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg":"user no exist", "refer":"/"}
	}else{
		password = com.Md5(password + user.Salt)
		if password == user.Password {
			this.SetSession("username", username)
			this.Data["json"] = map[string]interface{}{"result": true, "msg": "user login success ", "refer": "/admin"}
		}else{
			this.Data["json"] = map[string]interface{}{"result":false, "msg":"login failed", "refer":"/"}
		}
	}
	this.ServeJson()
}

//退出登陆
type LogoutController struct {
	controllers.BaseController
}
func(this *LogoutController)Get(){
	username := this.GetSession("username")
	if username == "" {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid params ", "refer": "/"}
		this.ServeJson()
	}
	this.DelSession("username")
	this.Ctx.WriteString("logout success")
}
func (this *LogoutController)Post(){
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request ", "refer": "/"}
	this.ServeJson()
}