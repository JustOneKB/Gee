package middlewares

import (
	"fmt"
	"gee"
	"time"
)

func InitMiddleware(c *gee.Context) {
	start := time.Now().UnixMicro()

	// c.Set("username", "张三")
	c.Next()
	end := time.Now().UnixMicro()
	fmt.Println("响应时间:", end-start)

}
