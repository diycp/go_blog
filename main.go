package main

import (
	"github.com/astaxie/beego"
	_ "blog/routers"
	_ "blog/initial"
)
func main(){
	beego.Run()
}