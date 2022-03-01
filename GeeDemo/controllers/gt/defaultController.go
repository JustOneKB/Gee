package controllers

import (
	"gee"
)

type DefaultControllers struct{}

func (con DefaultControllers) Index(c *gee.Context) {
	c.JSON(200, gee.H{
		"msg": "首页",
	})
	// username, err := c.Get("username")
	// if err {
	// 	fmt.Println(username)
	// } else {
	// 	fmt.Println("没收到中间件的数据")
	// }

	//cookie
	// c.SetCookie("username", "李四", 3600, "/", "localhost", false, true) //gee 暂未实现

	//设置session
	// session := sessions.Default(c) //session中间件中调用gin.Context, gee.Context不行 暂时注释
	// session.Options(sessions.Options{
	// 	MaxAge: 3600 * 6, //6h
	// })
	// session.Set("username", "王二 222")
	// session.Save() //设置结束之后必须保存

	// c.HTML(http.StatusOK, "default/index.html", gee.H{
	// 	"title": "这是首页",
	// 	"date":  1645443553,
	// })
}

func (con DefaultControllers) News(c *gee.Context) {
	//获取session
	// session := sessions.Default(c)

	// username := session.Get("username")
	// c.String(200, "session = %v", username)
}

/*
func (con DefaultControllers) Shop(c *gee.Context) {
	//获取cookie
	username, _ := c.Cookie("username") //gee暂未实现

	c.String(200, "cookie = "+username)
}
*/
