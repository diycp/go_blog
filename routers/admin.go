package routers
import (
	"github.com/astaxie/beego"
	"blog/controllers/admin"
)
func init(){
	beego.Router("/register", &admin.RegisterController{})
	beego.Router("/login", &admin.LoginController{})
	beego.Router("/logout", &admin.LogoutController{})
}