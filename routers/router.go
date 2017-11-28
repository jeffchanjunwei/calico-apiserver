package routers

import (
	"github.com/jeffchanjunwei/calico-apiserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
<<<<<<< HEAD
        ns := beego.NewNamespace("/calico/v1",
		beego.NSNamespace)"/ippool",
=======
        ns := beego.NewNamespace("/paas-net/v1",
		beego.NSNamespace("/ippool",
>>>>>>> dev-cjw
			beego.NSInclude(
				&controllers.IppoolController{},
			),
		),
	)
	beego.AddNamespace(ns)

//    	beego.Router("/calico/ippool", &controllers.IppoolController{})
//    	beego.Router("/calico/bgppeer", &controllers.BgppeerController{})
//	beego.Router("/calico/hostendpoint", &controlers.HostendpointController{})
//	beego.Router("/calico/node", &controllers.NodeController{})
//	beego.Router("/calico/policy", &controllers.PolicyController{})
//	beego.Router("/calico/profile", &controllers.ProfileController{})
//	beego.Router("/calico/workloadendpoint", &controllers.WorkloadendpointController{})

}
