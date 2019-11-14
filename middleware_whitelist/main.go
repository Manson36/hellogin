package main

import "github.com/gin-gonic/gin"

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string {
			"127.0.1.1",
		}
		flag := false
		clientIP := c.ClientIP()//获取本机的ip

		for _, host := range ipList {
			if host == clientIP {
				flag = true
				break
			}
		}

		if !flag {
			c.String(401, "%s, is not in ipList", clientIP)
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()

	r.Use(IPAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})

	r.Run()
}
