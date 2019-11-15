package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" finding:"required" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" finding:"required gt=CheckIn" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()

	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(&b); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "ok", "booking": b})
	})
}
