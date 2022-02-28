package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger()) // global midlleware

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.Run(":9999")
}
