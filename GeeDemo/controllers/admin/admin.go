package controllers

import (
	"gee"
	"net/http"
)

type AdminControllers struct{}

func (con AdminControllers) AdminIndex(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"msg": "这是后台首页",
	})
}
