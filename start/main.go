package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { //这里是一个回调方法
		c.JSON(200, gin.H{
			"message": "pong", //map结构
		})
	})

	r.Run()//默认端口8080
}
