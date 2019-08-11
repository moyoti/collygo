package routers

import (
	"demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Include(&controllers.TestController{})
	//beego.Include(&controllers.Testormtable2Controller{})
	//beego.Include(&controllers.TestormtableController{})
	ns := beego.NewNamespace("/test",
		beego.NSNamespace("/test",
			beego.NSInclude(&controllers.TestController{}),
		),
		beego.NSNamespace("/testormtable",
			beego.NSInclude(&controllers.TestormtableController{}),
		),
		beego.NSNamespace("/testormtable2",
			beego.NSInclude(&controllers.Testormtable2Controller{}),
		),
	)
	beego.AddNamespace(ns)
}
