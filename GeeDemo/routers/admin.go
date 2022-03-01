package routers

import (
	controllers "example/GeeDemo/controllers/admin"
	"gee"
)

func AdminRoutersInit(r *gee.Engine) {
	adminRouters := r.Group("/admin")
	adminRouters.GET("/", controllers.AdminControllers{}.AdminIndex)
	adminRouters.GET("/user", controllers.UserControllers{}.Index)
	adminRouters.GET("/user/add", controllers.UserControllers{}.Add)
	// adminRouters.POST("/user/doUpload", controllers.UserControllers{}.DoUpload)

}
