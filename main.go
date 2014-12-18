package main

import (
	"github.com/astaxie/beego"
	_ "github.com/rtanwar/bee_test/routers"
)

func main() {
	// beego.SetStaticPath("/static", "public")
	beego.Run()
}
