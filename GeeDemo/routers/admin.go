package routers

import (
	controllers "example/GeeDemo/controllers/admin"
	"gee"
)

func AdminRoutersInit(r *gee.Engine) {
	adminRouters := r.Group("/admin")
	adminRouters.GET("/", controllers.AdminIndex)
	adminRouters.GET("/userlist", controllers.AdminUserlist)
}
