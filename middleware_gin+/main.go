package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	//这里将log输出到文件，即请求的日志。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New() //这里不使用Default是因为，Default中已经实现了Logger和Recovery两个中间件
	r.Use(gin.Logger())
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		c.String(200, "%s", name)
	})
	r.Run()
}
