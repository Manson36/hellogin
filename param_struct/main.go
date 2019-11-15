package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

//参数绑定到结构体的功能
type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	BirthDay time.Time `form:"birthday" time_format:"2006-01-02"` //后面的tag是设置time的格式
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)

	//测试一下效果：curl -X GET "http://127.0.0.1:8080/testing?name=wang&address=wuhan&birthday=2006-01-06"
	//====》这里出现了问题？？？？
	r.Run()
}

func testing(c *gin.Context) {
	var person Person
	//这里根据请求content-type 来做不同的binding操作
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(200, "error:%v", err)
	}
}
