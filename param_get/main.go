package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//请求路由-获取请求参数
func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		//这里我们两个参数，一个是没有默认值，一个有默认值
		firstName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "last_default_name")

		c.String(http.StatusOK, "%s, %s", firstName, lastName)
	})

	//我们测试一下效果：curl -X GET 'http://127.0.0.1:8080/test?first_name=wang'
	//				curl -X GET 'http://127.0.0.1:8080/test?first_name=wang&last_name=kai'
	r.Run(":8080")
}
