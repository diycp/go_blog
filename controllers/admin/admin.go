package admin
import "blog/controllers"
type AdminController struct {
	controllers.AdminBaseController
}
func (this *AdminController)Get(){
	this.TplNames = "admin/index.tpl"
}
func (this *AdminController)Post(){
	this.Data["json"] = map[string]interface{}{"result":false, "msg":"invalid request", "refer":"/"}
	this.ServeJson()
}