package controllers

import "webless/controllers/base"

// AsZoneObjectInfoController operations for AsZoneObjectInfo
type AsBaseController struct {
	base.BaseController
}

// URLMapping ...
func (c *AsBaseController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title Get All
// @Description get []Stalk
// @Success 200 {object} base.Stalk
// @Failure 403
// @router / [get]
func (c *AsBaseController) GetAll() {
	defer func() {
		if e := recover(); e != nil {
			err := e.(error)
			c.Data["json"] = err.Error()
		}
	}()
	v := base.GetTree()
	c.Data["json"] = v
}
