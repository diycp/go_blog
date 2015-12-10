package main

import (
//	"github.com/astaxie/beego"
	"fmt"
	"github.com/gogather/com"
)
func main(){
	content := `{"name":"lixuan", "age":"23"}`
	json, _ := com.JsonDecode(content)
	json = json.(map[string]interface{})["res"]
	fmt.Println(json)
//	if _, ok := json.(map[string]interface{}); !ok {
//		fmt.Println("json error")
//	}
	if _, ok := json.(map[string]interface{}); !ok != nil {
		fmt.Println("json error")
	}
	fmt.Println(123)
	fmt.Println(456)
//	beego.Run()
}