package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Age int 			`form:"age" binding:"required, gt=10"`//binding:是验证内容，“，”表示同时满足，“|”表示满足一个即可
	Name string 		`form:"name" binding:"required"`	  //gt=10,表示输入的值要大于10
	Address string		`form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	//功能与参数绑定结构体类似，这里需要前置条件是参数已经绑定到结构体，然后基于这个结构体再进行验证。
	//只是在结构体上增加了验证规则

	r.GET("/testing", func(c *gin.Context) {
		var person Person

		if err := c.ShouldBind(person); err != nil {
			c.String(500, "%v", err)
			c.Abort()
			return
		}
		c.String(200, "%v", person)
	})

	r.Run()
}
