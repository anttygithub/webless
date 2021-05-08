package base

import "net/http"

const (
	ActionKindBase     = "base"
	ActionKindAddition = "addition"
)

var ActionsMap = newActions()

// ActionsInit init
func newActions() (acts map[string][]Action) {
	acts = make(map[string][]Action)
	act := Action{}
	obj := ""
	// as_zone_object_info
	obj = "as_zone_object_info"
	// --------------------------
	act.Name = "add"
	act.Desc = "新增记录"
	act.Kind = ActionKindBase
	act.Method = http.MethodPost
	act.Primary = ""
	acts[obj] = append(acts[obj], act)
	// --------------------------
	act.Name = "get by id"
	act.Desc = "根据Id查询"
	act.Kind = ActionKindBase
	act.Method = http.MethodGet
	act.Primary = "Id"
	acts[obj] = append(acts[obj], act)
	// --------------------------
	act.Name = "get all"
	act.Desc = "查询所有记录"
	act.Kind = ActionKindBase
	act.Method = http.MethodGet
	act.Primary = ""
	acts[obj] = append(acts[obj], act)
	// --------------------------
	act.Name = "updata"
	act.Desc = "修改记录"
	act.Kind = ActionKindBase
	act.Method = http.MethodPut
	act.Primary = "Id"
	acts[obj] = append(acts[obj], act)
	// --------------------------
	act.Name = "delete"
	act.Desc = "删除记录"
	act.Kind = ActionKindBase
	act.Method = http.MethodDelete
	act.Primary = "Id"
	acts[obj] = append(acts[obj], act)

	// as_zone_object_option
	obj = "as_zone_object_option"
	// --------------------------
	act.Name = "add"
	act.Desc = "新增记录"
	act.Kind = ActionKindBase
	act.Method = http.MethodPost
	act.Primary = ""
	acts[obj] = append(acts[obj], act)
	// --------------------------
	act.Name = "get by id"
	act.Desc = "根据Id查询"
	act.Kind = ActionKindBase
	act.Method = http.MethodGet
	act.Primary = "Id"
	acts[obj] = append(acts[obj], act)
	// --------------------------
	act.Name = "get all"
	act.Desc = "查询所有记录"
	act.Kind = ActionKindBase
	act.Method = http.MethodGet
	act.Primary = ""
	acts[obj] = append(acts[obj], act)
	// --------------------------
	act.Name = "updata"
	act.Desc = "修改记录"
	act.Kind = ActionKindBase
	act.Method = http.MethodPut
	act.Primary = "Id"
	acts[obj] = append(acts[obj], act)
	// --------------------------
	act.Name = "delete"
	act.Desc = "删除记录"
	act.Kind = ActionKindBase
	act.Method = http.MethodDelete
	act.Primary = "Id"
	acts[obj] = append(acts[obj], act)

	// -------------------------- end line --------------------------

	return
}
