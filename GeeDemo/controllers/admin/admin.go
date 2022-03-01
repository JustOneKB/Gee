package controllers

import "gee"

func AdminIndex(c *gee.Context) {
	c.JSON(200, gee.H{
		"msg": "这是后台首页",
	})
}

func AdminUserlist(c *gee.Context) {
	c.JSON(200, gee.H{
		"msg": "后台---用户名单",
	})
}
