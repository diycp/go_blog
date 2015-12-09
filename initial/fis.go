package initial
import (
	"github.com/astaxie/beego"
	"github.com/gogather/com"
	"html/template"
)

func loadMap() string{
	mapPath := beego.AppConfig.String("static_map")
	mapContent := com.ReadFile(mapPath)
	return mapContent
}

func Fis(key string) template.HTML {
	runmode := beego.AppConfig.String("RunMode")
//	if runmode == "dev" {
//		text := ""
//		uri := "/static/"
//	}
}