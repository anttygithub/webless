package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/client/orm"
)

type AsZoneObjectInfo struct {
	Id                int       `kind:"pk" orm:"pk;column(id);auto"`
	Name              string    `kind:"base" orm:"column(name);size(200)" description:"名称"`
	PrimaryKey        string    `kind:"base" orm:"column(primary_key);size(200)" description:"主键"`
	RefAttribute      string    `kind:"base" orm:"column(ref_attribute);null" description:"关联信息"`
	BasicAttribute    string    `kind:"base" orm:"column(basic_attribute);null" description:"基础信息"`
	AdditionAttribute string    `kind:"base" orm:"column(addition_attribute);null" description:"附加信息"`
	Comment           string    `kind:"base" orm:"column(comment);size(500);null" description:"说明"`
	IsActive          int8      `kind:"base" orm:"column(is_active)" description:"是否启用,1启用0不启用,默认不启用"`
	AddUser           string    `kind:"addition" orm:"column(add_user);size(128);null"`
	UpdateUser        string    `kind:"addition" orm:"column(update_user);size(128);null"`
	CreateTm          time.Time `kind:"addition" orm:"column(create_tm);type(datetime);auto_now_add"`
	UpdateTm          time.Time `kind:"addition" orm:"column(update_tm);type(datetime);auto_now"`
}

func (t *AsZoneObjectInfo) TableName() string {
	return "as_zone_object_info"
}

func init() {
	orm.RegisterModel(new(AsZoneObjectInfo))
}

// AddAsZoneObjectInfo insert a new AsZoneObjectInfo into database and returns
// last inserted Id on success.
func AddAsZoneObjectInfo(m *AsZoneObjectInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAsZoneObjectInfoById retrieves AsZoneObjectInfo by Id. Returns error if
// Id doesn't exist
func GetAsZoneObjectInfoById(id int) (v *AsZoneObjectInfo, err error) {
	o := orm.NewOrm()
	v = &AsZoneObjectInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAsZoneObjectInfo retrieves all AsZoneObjectInfo matches certain condition. Returns empty list if
// no records exist
func GetAllAsZoneObjectInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AsZoneObjectInfo))
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

	var l []AsZoneObjectInfo
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

// UpdateAsZoneObjectInfo updates AsZoneObjectInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateAsZoneObjectInfoById(m *AsZoneObjectInfo) (err error) {
	o := orm.NewOrm()
	v := AsZoneObjectInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAsZoneObjectInfo deletes AsZoneObjectInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAsZoneObjectInfo(id int) (err error) {
	o := orm.NewOrm()
	v := AsZoneObjectInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AsZoneObjectInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
