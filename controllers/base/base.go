package base

import (
	"reflect"
	"strings"
	"time"

	beego "github.com/astaxie/beego/server/web"
)

const (
	//CollectionType type of collection data
	CollectionType = "Collection"
	//ErrorType type of error message
	ErrorType     = "Error"
	MsgNoRow      = "get 0 record"
	LevelPk       = "pk"
	LevelBase     = "base"
	LevelAddition = "addition"
)

//StandardOutput all the api use this to make standard api format
type StandardOutput struct {
	Timestamp    int64           `json:"timestamp"`
	ObjectName   string          `json:"objectName"`
	Msg          string          `json:"msg"`
	Data         interface{}     `json:"data"`
	ErrorCode    int             `json:"errorCode"`
	ErrorMessage string          `json:"errorMessage"`
	Total        int             `json:"total,omitempty"`
	CurrentPage  int             `json:"currentPage,omitempty"`
	CountPerPage int             `json:"countPerPage,omitempty"`
	Actions      []Action        `json:"actions,omitempty"`
	Attributes   []AttributeDesc `json:"attributes,omitempty"`
}

//Attributes all the api use this to make standard Attributes format
// type Attributes struct {
// 	Primary            string            `json:"primary"`
// 	BasicAttributes    string            `json:"basic_attributes"`
// 	AdditionAttributes string            `json:"addition_attributes"`
// 	RefAttributes      string            `json:"ref_attributes"`
// 	HeaderList         map[string]string `json:"header_list"`
// }

// AttributeDesc Attribute description
type AttributeDesc struct {
	Index    int    `json:"index"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Kind     string `json:"kind"`
	Relation string `json:"relation_attribute"`
}

//Action all the api use this to make standard Action format
type Action struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Kind    string `json:"kind"`
	Method  string `json:"method"`
	URI     string `json:"uri"`
	Primary string `json:"primary"`
	Auth    string `json:"auth"`
}

//BaseController base controller struct
type BaseController struct {
	beego.Controller
}

//Finish return data is written here
func (b *BaseController) Finish() {
	so := pack(b.Data["json"], b.GetObjectName())
	if so.ErrorCode != 0 {
		b.Ctx.ResponseWriter.WriteHeader(500)
	}
	b.Data["json"] = &so
	b.ServeJSON()
}

//GetType return type of request
func (b *BaseController) GetObjectName() string {
	u := strings.Split(b.Ctx.Request.URL.Path, "/")
	if len(u) >= 2 {
		return u[2]
	} else {
		return ""
	}
}

func pack(data interface{}, obj_name string) *StandardOutput {
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}
	var so *StandardOutput
	switch dataValue.Kind() {
	case reflect.Slice:
		so = packSlice(data, obj_name, dataValue.Len())
	case reflect.Struct:
		so = packStruct(data, dataValue.Type(), obj_name)
	case reflect.String:
		so = packString(dataValue.String(), obj_name)
	default:
		so = packStruct(data, dataValue.Type(), obj_name)
	}
	return so
}

func packSlice(data interface{}, obj_name string, len int) *StandardOutput {
	if len == 0 {
		return &StandardOutput{Timestamp: time.Now().Unix(), ObjectName: obj_name + "_list", Total: 0, Data: []string{}, Actions: GetActions(obj_name), Msg: MsgNoRow}
	} else {
		d := reflect.ValueOf(data)
		dataValue := reflect.ValueOf(d.Index(0).Interface())
		if dataValue.Kind() == reflect.Ptr {
			dataValue = dataValue.Elem()
		}
		return &StandardOutput{Timestamp: time.Now().Unix(), ObjectName: obj_name + "_list", Total: len, Attributes: GetAttributes(dataValue.Type()), Data: data, Actions: GetActions(obj_name)}
	}
}

func packStruct(data interface{}, model_type reflect.Type, obj_name string) *StandardOutput {
	return &StandardOutput{Timestamp: time.Now().Unix(), ObjectName: obj_name, Attributes: GetAttributes(model_type), Data: data, Actions: GetActions(obj_name)}
}

func packString(msg string, obj_name string) (so *StandardOutput) {
	if msg != "ok" {
		return &StandardOutput{Timestamp: time.Now().Unix(), ObjectName: obj_name, ErrorCode: 1, ErrorMessage: msg}
	} else {
		return &StandardOutput{Timestamp: time.Now().Unix(), ObjectName: obj_name, Msg: msg, Actions: GetActions(obj_name)}
	}
}

//GetAttributes get object's Attribute
func GetReflectType(v reflect.Value) reflect.Type {
	for i := 0; i < 10; i++ {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		} else {
			return v.Type()
		}
	}
	return v.Type()
}

//GetAttributes get object's Attribute
func GetAttributes(t reflect.Type) (attrs []AttributeDesc) {
	if t == nil {
		return
	}
	for i := 0; i < t.NumField(); i++ {
		dfield := t.Field(i)
		attr := AttributeDesc{
			Name:     dfield.Name,
			Index:    i + 1,
			Desc:     dfield.Tag.Get("description"),
			Kind:     dfield.Tag.Get("kind"),
			Relation: dfield.Tag.Get("ref"),
		}
		attrs = append(attrs, attr)
	}
	return
}

//BaseController base controller struct
type BaseControllerAction interface {
	Actions() []string
}

//getActions get object's actions
func GetActions(obj_name string) (acts []Action) {
	apiversion, _ := beego.AppConfig.String("apiversion")
	if _, ok := ActionsMap[obj_name]; !ok {
		return
	}
	acts = ActionsMap[obj_name]
	for i := 0; i < len(acts); i++ {
		acts[i].URI = "/" + apiversion + "/" + obj_name
		if acts[i].Primary != "" {
			acts[i].URI += "/:" + acts[i].Primary
		}
	}
	return
}
