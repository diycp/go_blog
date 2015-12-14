package controllers
import (
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController)Forbidden(mark, condition string){
	mark = strings.ToLower(mark)
	condition = strings.ToLower(condition)
	if makr == "not" && this.Data["userIs"] != condition {
		this.Redirect("/", 302)
	}else if this.Data["userIs"] == condition {
		this.Redirect("/", 302)
	}
}

func (this *BaseController)Prepare(){
	user := this.GetSession("username")
	if user == nil {
		this.Data["userIs"] = ""
	}else{
		this.Data["userIs"] = "admin"
	}
	this.Data["isDev"] = beego.AppConfig.String("RunMode") == "dev"
}