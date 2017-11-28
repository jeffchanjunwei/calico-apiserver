package main

import (
	"flag"   

	_ "github.com/jeffchanjunwei/calico-apiserver/routers"
        _ "github.com/jeffchanjunwei/calico-apiserver/controllers"
	"github.com/astaxie/beego"
)

func main() {
	flag.Parse()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/docs"] = "swagger"
	}

	beego.Run()
}

