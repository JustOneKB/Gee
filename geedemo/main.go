package main

import (
	"gee"
	"geedemo/routers"
)

func main() {
	r := gee.New()

	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)

	r.Run(":8000")
}
