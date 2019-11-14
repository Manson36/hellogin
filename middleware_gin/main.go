package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()//这里不使用Default是因为，Default中已经实现了Logger和Recovery两个中间件
	//这里的输出内容是基于控制台输出
	r.Use(gin.Logger())
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		c.String(200, "%s", name)
	})
	r.Run()
}
