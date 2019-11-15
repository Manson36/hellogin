package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//路由请求-获取body内容
func main() {
	r := gin.Default()

	r.POST("/test", func(c *gin.Context) {
		//我们要获取body的内容，还要学习ioutil.ReadAll方法
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort() //用来把输出结束
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		//c.String(http.StatusOK, string(bodyBytes)) //简单测试readAll效果
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default_last_name")

		c.String(http.StatusOK, "%s, %s, %s", firstName, lastName, string(bodyBytes))
		//这里我们调用完readAll之后是无法PostForm的数值是拿不到的，怎么解决这个问题：
		//在readAll之后加上代码：（目的将内容存到request.body中）
		//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	})

	//测试效果：//curl -X POST "http://127.0.0.1:8080/test" -d "{"name": "wang"}"
	//		curl -X POST "http://127.0.0.1:8080/test" -d "first_name=wang&last_name=kai"
	r.Run()
}
