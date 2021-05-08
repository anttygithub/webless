package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/client/orm"
)

type AsZoneObjectOption struct {
	Id         int       `kind:"pk" orm:"column(id);auto"`
	ObjId      int       `kind:"base" ref:"as_zone_object_info.Id" orm:"column(obj_id)" description:"对象id"`
	OptName    string    `kind:"base" orm:"column(opt_name);size(200)" description:"操作名"`
	Protocol   string    `kind:"base" orm:"column(protocol);size(200)" description:"协议"`
	Ip         string    `kind:"base" orm:"column(ip);size(50)" description:"服务地址"`
	Port       string    `kind:"base" orm:"column(port);size(10)" description:"服务端口"`
	Uri        string    `kind:"base" orm:"column(uri);size(200)" description:"URI"`
	Method     string    `kind:"base" orm:"column(method);size(50)" description:"方法（POST,GET,PUT,DELETE......）"`
	Comment    string    `kind:"base" orm:"column(comment);size(500);null" description:"说明"`
	IsActive   int8      `kind:"base" orm:"column(is_active)" description:"是否启用,1启用0不启用,默认不启用"`
	AddUser    string    `kind:"addition" orm:"column(add_user);size(128);null"`
	UpdateUser string    `kind:"addition" orm:"column(update_user);size(128);null"`
	CreateTm   time.Time `kind:"addition" orm:"column(create_tm);type(datetime);auto_now_add"`
	UpdateTm   time.Time `kind:"addition" orm:"column(update_tm);type(datetime);auto_now"`
}

func (t *AsZoneObjectOption) TableName() string {
	return "as_zone_object_option"
}

func init() {
	orm.RegisterModel(new(AsZoneObjectOption))
}

// AddAsZoneObjectOption insert a new AsZoneObjectOption into database and returns
// last inserted Id on success.
func AddAsZoneObjectOption(m *AsZoneObjectOption) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAsZoneObjectOptionById retrieves AsZoneObjectOption by Id. Returns error if
// Id doesn't exist
func GetAsZoneObjectOptionById(id int) (v *AsZoneObjectOption, err error) {
	o := orm.NewOrm()
	v = &AsZoneObjectOption{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAsZoneObjectOption retrieves all AsZoneObjectOption matches certain condition. Returns empty list if
// no records exist
func GetAllAsZoneObjectOption(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AsZoneObjectOption))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []AsZoneObjectOption
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateAsZoneObjectOption updates AsZoneObjectOption by Id and returns error if
// the record to be updated doesn't exist
func UpdateAsZoneObjectOptionById(m *AsZoneObjectOption) (err error) {
	o := orm.NewOrm()
	v := AsZoneObjectOption{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAsZoneObjectOption deletes AsZoneObjectOption by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAsZoneObjectOption(id int) (err error) {
	o := orm.NewOrm()
	v := AsZoneObjectOption{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AsZoneObjectOption{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
