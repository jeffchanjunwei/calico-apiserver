package main

import (
	_ "github.com/jeffchanjunwei/calico-apiserver/routers"
        _ "github.com/jeffchanjunwei/calico-apiserver/controllers"
	"github.com/projectcalico/libcalico-go/lib/client"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	flag.Parse()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/docs"] = "swagger"
	}

	// calioClient, err := client.NewFromEnv()
        // if err != null {
        //        logs.Error("Error calico client initialization.")
        //}
	beego.Run()
}

