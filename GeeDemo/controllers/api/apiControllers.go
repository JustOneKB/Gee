package controllers

import "gee"

type ApiControllers struct{}

func (con ApiControllers) ApiIndex(c *gee.Context) {
	c.JSON(200, gee.H{
		"msg": "这是接口首页",
	})
}

func (con ApiControllers) ApiAdd(c *gee.Context) {
	c.JSON(200, gee.H{
		"msg": "增加一个接口",
	})
}

func (con ApiControllers) ApiDelete(c *gee.Context) {
	c.JSON(200, gee.H{
		"msg": "删除一个接口",
	})
}
