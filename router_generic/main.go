package main

import "github.com/gin-gonic/gin"

//请求路由-泛绑定前缀
func main() {
	r := gin.Default()

	//这里我们泛绑定的是user这个前缀的，下面是*action，也就是说所有的user前缀都打到这一个相同的回调函数中
	r.GET("/user/*action", func(c *gin.Context) {
			c.String(200, "hello world")
	})

	//测试一下效果：curl -X GET "http://127.0.0.1:8080/user/xxx111"
	r.Run()
}
