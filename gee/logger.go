package gee

import (
	"log"
	"time"
)

//一种中间件举例
func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()

		c.Next()

		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
