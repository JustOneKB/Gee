package main

import (
	"example/GeeDemo/routers"
	"gee"
)

func main() {
	r := gee.New()

	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)

	r.Run(":8000")
}
