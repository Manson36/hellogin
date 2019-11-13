package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//请求路由-绑定静态文件夹
func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")//前面设置是路由，后面是相对文件夹
	//另外一种写法
	r.StaticFS("/static", http.Dir("static"))//前面设置是路由，后面是它的资源
	//另一种，设置单个静态文件
	r.StaticFile("/favicon.ico", "./favicon.ico")//前面设置是路由，后面是它的资源

	//在启动之前，我们要添加文件夹，assets,里面添加html文件
	r.Run()
}
