package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
)

//这里我们要加一个特殊的处理：我们预约要比今天的时间还要大才行
//valid v8支持自定义结构体，我们期望定义一个validator标签，然后针对标签定义一下验证方法，然后把这个方法注册到validator验证器里边，完成测试
//自定义验证器可以在validator-goDoc中查找，需要引入validator的内库

type Booking struct {
	CheckIn  time.Time `form:"check_in" finding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" finding:"required gt=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if data, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if data.Unix() > today.Unix() {
			return true
		}
	}

	return false
}

func main() {
	r := gin.Default()

	//将方法注册到validator里面中
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(&b); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "ok", "booking": b})
	})

	r.Run()
}
