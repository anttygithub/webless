// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"webless/controllers"

	beego "github.com/astaxie/beego/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/stalk",
			beego.NSInclude(
				&controllers.AsBaseController{},
			),
		),

		beego.NSNamespace("/as_zone_object_info",
			beego.NSInclude(
				&controllers.AsZoneObjectInfoController{},
			),
		),

		beego.NSNamespace("/as_zone_object_option",
			beego.NSInclude(
				&controllers.AsZoneObjectOptionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
