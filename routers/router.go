package routers

import (
	"github.com/astaxie/beego"
	"goyo.in/gpstracker/controllers"
)

func init() {

	// namespaces
	var namespaces []string = []string{"goyoapi", "another"}

	// controllers
	var ctrllers []string = []string{"login", "other"}

	restfulRouter := beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/"+ctrllers[0],
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
	)
	beego.AddNamespace(restfulRouter)

	restfulRouter = beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/getvahicleupdates",
			beego.NSInclude(
				&controllers.TripDataController{},
			),
		),
	)
	beego.AddNamespace(restfulRouter)
}
