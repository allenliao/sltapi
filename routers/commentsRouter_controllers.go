package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["sltapi/controllers:GameLoginController"] = append(beego.GlobalControllerRouter["sltapi/controllers:GameLoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["sltapi/controllers:GameLoginController"] = append(beego.GlobalControllerRouter["sltapi/controllers:GameLoginController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sltapi/controllers:GameLoginController"] = append(beego.GlobalControllerRouter["sltapi/controllers:GameLoginController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sltapi/controllers:GameLoginController"] = append(beego.GlobalControllerRouter["sltapi/controllers:GameLoginController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["sltapi/controllers:GameLoginController"] = append(beego.GlobalControllerRouter["sltapi/controllers:GameLoginController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
