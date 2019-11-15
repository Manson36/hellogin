package main

import "github.com/gin-gonic/gin"

//模板渲染
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*") //这个方法可以将html文件倒入

	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "index.html",
		})
	})
	//这里我们使用的是goland启动，默认在hellogin文件夹，需要跳转到other_template文件执行
	//curl -X GET "http://127.0.0.1:8080/index"
	r.Run()
}
