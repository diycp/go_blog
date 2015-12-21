package admin
import (
	"blog/controllers"
	"github.com/astaxie/beego"
	"github.com/duguying/blog/utils"
	"blog/models"
	"fmt"
)
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

	if !utils.CheckUsername(username){
		this.Data["json"] = map[string]interface{}{"result":false, "msg":"illegal username", "refer":"/"}
		this.ServeJson()
		return
	}
	id, err := models.AddUser(username, password)
	if err == nil && id{
		this.Data["json"] = map[string]interface{}{"result":true, "msg":"register success", "refer":"/"}
	}else{
		this.Data["json"] = map[string]interface{}{"result": false, "msg":"register failed", "refer":"/"}
	}
	this.ServeJson()
}