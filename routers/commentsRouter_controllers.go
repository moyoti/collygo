package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["demo/controllers:TestController"] = append(beego.GlobalControllerRouter["demo/controllers:TestController"],
		beego.ControllerComments{
			Method:           "CollectData",
			Router:           `/col`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:TestController"] = append(beego.GlobalControllerRouter["demo/controllers:TestController"],
		beego.ControllerComments{
			Method:           "DownloadFile",
			Router:           `/download`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:TestController"] = append(beego.GlobalControllerRouter["demo/controllers:TestController"],
		beego.ControllerComments{
			Method:           "Test",
			Router:           `/tt`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:Testormtable2Controller"] = append(beego.GlobalControllerRouter["demo/controllers:Testormtable2Controller"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:Testormtable2Controller"] = append(beego.GlobalControllerRouter["demo/controllers:Testormtable2Controller"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:Testormtable2Controller"] = append(beego.GlobalControllerRouter["demo/controllers:Testormtable2Controller"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:Testormtable2Controller"] = append(beego.GlobalControllerRouter["demo/controllers:Testormtable2Controller"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:Testormtable2Controller"] = append(beego.GlobalControllerRouter["demo/controllers:Testormtable2Controller"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:TestormtableController"] = append(beego.GlobalControllerRouter["demo/controllers:TestormtableController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:TestormtableController"] = append(beego.GlobalControllerRouter["demo/controllers:TestormtableController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:TestormtableController"] = append(beego.GlobalControllerRouter["demo/controllers:TestormtableController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:TestormtableController"] = append(beego.GlobalControllerRouter["demo/controllers:TestormtableController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["demo/controllers:TestormtableController"] = append(beego.GlobalControllerRouter["demo/controllers:TestormtableController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
