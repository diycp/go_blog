package controllers
import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController)Error404(){
	this.TplNames = "error/404.tpl"
}

func (this *ErrorController)Error501(){
	this.TplNames = "error/501.tpl"
}