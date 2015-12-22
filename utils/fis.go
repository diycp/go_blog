package utils
import (
	"github.com/astaxie/beego"
	"github.com/gogather/com"
	"html/template"
	"strings"
)

func loadMap() string{
	mapPath := beego.AppConfig.String("static_map")
	mapContent := com.ReadFile(mapPath)
	return mapContent
}

func Fis(key string) template.HTML {
	runMode := beego.AppConfig.String("RunMode")
	text := ""
	//如果是开发环境
	if runMode == "dev" {
		uri := "/static/" + key
		if strings.HasSuffix(uri, "css"){
			text = `<link rel="stylesheet" href="` + uri + `">`
		}else if strings.HasSuffix(uri, "js"){
			text = `<script src="` + uri + `"></script>`
		}
		return template.HTML(text)
	}
	//如果不是开发环境
	content := loadMap()
	json, ok := com.JsonDecode(content)
	if ok != nil {
		Exit(1, "json failed")
	}
	json = json.(map[string]interface{})["res"]
	if fileMap, ok := json.(map[string]interface{}); !ok {
		Exit(2, "map.json id illegal")
	}else{
		for tmpKey, views := range fileMap {
			uri, ok := views.(map[string]interface{})["uri"].(string)
			if !ok {
				Exit(3, "error in map.json")
			}
			fileType, ok := views.(map[string]interface{})["type"].(string)
			if !ok {
				Exit(4, "error in map.json")
			}
			if tmpKey == key {
				if fileType == "css" {
					text = `<link rel="stylesheet" href="` + uri + `">`
				}else if fileType == "js" {
					text = `<script src="` + uri + `"></script>`
				}
			}
		}
	}
	return template.HTML(text)
}