package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"] = append(beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"] = append(beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"] = append(beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:ippool`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"] = append(beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:ippool`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"] = append(beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"],
		beego.ControllerComments{
			Method: "Apply",
			Router: `/:ippool`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"] = append(beego.GlobalControllerRouter["github.com/jeffchanjunwei/calico-apiserver/controllers:IppoolController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:ippool`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
