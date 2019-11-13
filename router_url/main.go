package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	//设置一个get的路由，然后里面设置一些获取name的rul参数，获取id的rul参数，然后设置它的回调方法
	r.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"name": c.Param("name"),
		"id": c.Param("id"),
		})
	})

	//运行一下：curl "http://127.0.0.1:8080/lihua/12"
	r.Run()
	}
