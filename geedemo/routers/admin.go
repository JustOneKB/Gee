package routers

import (
	"gee"
	controllers "geedemo/controllers/admin"
)

func AdminRoutersInit(r *gee.Engine) {
	adminRouters := r.Group("/admin")
	adminRouters.GET("/", controllers.AdminIndex)
	adminRouters.GET("/userlist", controllers.AdminUserlist)
}
