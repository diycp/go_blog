package initial
import (
	"github.com/astaxie/beego"
)

func InitEnv(){
	runMode := beego.AppConfig.String("RunMode")
	if runMode == "dev" {
		beego.SetStaticPath("/static", "static/theme/default")
	}
}