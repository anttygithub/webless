package routers

import (
	beego "github.com/astaxie/beego/server/web"
	"github.com/astaxie/beego/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["webless/controllers:AsBaseController"] = append(beego.GlobalControllerRouter["webless/controllers:AsBaseController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webless/controllers:AsZoneObjectInfoController"] = append(beego.GlobalControllerRouter["webless/controllers:AsZoneObjectInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webless/controllers:AsZoneObjectInfoController"] = append(beego.GlobalControllerRouter["webless/controllers:AsZoneObjectInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webless/controllers:AsZoneObjectInfoController"] = append(beego.GlobalControllerRouter["webless/controllers:AsZoneObjectInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webless/controllers:AsZoneObjectInfoController"] = append(beego.GlobalControllerRouter["webless/controllers:AsZoneObjectInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webless/controllers:AsZoneObjectInfoController"] = append(beego.GlobalControllerRouter["webless/controllers:AsZoneObjectInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webless/controllers:AsZoneObjectOptionController"] = append(beego.GlobalControllerRouter["webless/controllers:AsZoneObjectOptionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webless/controllers:AsZoneObjectOptionController"] = append(beego.GlobalControllerRouter["webless/controllers:AsZoneObjectOptionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webless/controllers:AsZoneObjectOptionController"] = append(beego.GlobalControllerRouter["webless/controllers:AsZoneObjectOptionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webless/controllers:AsZoneObjectOptionController"] = append(beego.GlobalControllerRouter["webless/controllers:AsZoneObjectOptionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webless/controllers:AsZoneObjectOptionController"] = append(beego.GlobalControllerRouter["webless/controllers:AsZoneObjectOptionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
