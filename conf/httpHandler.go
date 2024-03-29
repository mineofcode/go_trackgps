package conf

import (
	"github.com/astaxie/beego"
	//load router
	_ "goyo.in/gpstracker/routers"
	_ "goyo.in/gpstracker/socketios"
)

func RestfulAPIServiceInit(method string) {

	/*beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
	  if ctx.Input.Method() == "OPTIONS" {
	    ctx.Output.Header("Access-Control-Allow-Origin", "*")
	    ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT")
	    ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	    ctx.Abort(200, "Error in POST - GET Method")
	  }
	})*/
	beego.BConfig.RunMode = "dev"
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.WebConfig.EnableDocs = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.CopyRequestBody = true
	//   beego.BConfig.Listen.HTTPSCertFile = "tls-ssl/file-rest.crt"
	//   beego.BConfig.Listen.HTTPSKeyFile = "tls-ssl/file-rest.key"
	beego.BConfig.Listen.HTTPPort = 6979
	beego.BConfig.Listen.HTTPSPort = 8081
	if method == "HTTP" {
		beego.BConfig.Listen.EnableHTTP = true
		beego.BConfig.Listen.EnableHTTPS = false
	} else if method == "HTTPS" {
		beego.BConfig.Listen.EnableHTTP = false
		beego.BConfig.Listen.EnableHTTPS = true
	}
	beego.Run()
}
