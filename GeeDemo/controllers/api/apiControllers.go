package controllers

import "gee"

func ApiIndex(c *gee.Context) {
	c.JSON(200, gee.H{
		"msg": "这是接口首页",
	})
}

func ApiAdd(c *gee.Context) {
	c.JSON(200, gee.H{
		"msg": "增加一个接口",
	})
}

func ApiDelete(c *gee.Context) {
	c.JSON(200, gee.H{
		"msg": "删除一个接口",
	})
}
