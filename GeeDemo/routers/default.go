package routers

import (
	controllers "example/GeeDemo/controllers/gt"
	"example/GeeDemo/middlewares"
	"gee"
)

func DefaultRoutersInit(r *gee.Engine) {
	defaultRouters := r.Group("/")
	defaultRouters.Use(middlewares.InitMiddleware)
	defaultRouters.GET("/", controllers.DefaultControllers{}.Index)
	// defaultRouters.GET("/news", controllers.DefaultControllers{}.News) //
	// defaultRouters.GET("/shop", controllers.DefaultControllers{}.Shop)
}
