package main

import (
	_ "webless/routers"

	"github.com/astaxie/beego/client/orm"
	beego "github.com/astaxie/beego/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sqlconn, _ := beego.AppConfig.String("sqlconn")
	orm.RegisterDataBase("default", "mysql", sqlconn)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

// func test() {
// 	test_data := testmodel{a: "1", b: "2"}
// 	logrus.Info("abctest------start-------")
// 	var ba BaseControllerAction = test_data
// 	logrus.Info(ba.Actions())
// 	logrus.Info(tt(test_data))
// 	logrus.Info("abctest-----finished-----")
// }

// type BaseControllerAction interface {
// 	Actions() []string
// }

// type testmodel struct {
// 	a string
// 	b string
// }

// func (this testmodel) Actions() []string {
// 	return []string{"add", "delete:id", "update"}
// }

// func GetActions(ba BaseControllerAction) (actions map[string]string) {
// 	actions = make(map[string]string)
// 	val := reflect.ValueOf(ba)
// 	typ := reflect.Indirect(val).Type()
// 	act_list := ba.Actions()
// 	for i := 0; i < len(act_list); i++ {
// 		typ_arr := strings.Split(fmt.Sprint(typ), ".")
// 		typ_str := fmt.Sprint(typ)
// 		if len(typ_arr) > 0 {
// 			typ_str = typ_arr[len(typ_arr)-1]
// 		}
// 		key_arr := strings.Split(act_list[i], ":")
// 		actions[key_arr[0]] = "/v1/" + typ_str + "/" + key_arr[0]
// 		if len(key_arr) > 1 {
// 			actions[key_arr[0]] += "/:" + key_arr[1]
// 		}
// 	}
// 	return actions
// }
