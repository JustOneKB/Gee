package routers

import (
	"gee"
	controllers "geedemo/controllers/api"
)

/*
	这个包主要存放不同group的路由声明，
	那么主函数中只需要调用不同分组声明的函数就可以完成路由的注册，
*/
func ApiRoutersInit(r *gee.Engine) {
	apiRouters := r.Group("/api")
	apiRouters.GET("/", controllers.ApiIndex)
	apiRouters.GET("/add", controllers.ApiAdd)
	apiRouters.GET("/delete", controllers.ApiDelete)
}
