package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//请求路由-绑定静态文件夹
func main() {
	r := gin.Default()

	r.Static("/assets", "./assets") //前面设置是路由，后面是相对文件夹
	//另外一种写法
	r.StaticFS("/static", http.Dir("static")) //前面设置是路由，后面是它的资源
	//另一种，设置单个静态文件
	r.StaticFile("/favicon.ico", "./favicon.ico") //前面设置是路由，后面是它的资源

	//在启动之前，我们要添加文件夹，assets,里面添加html文件
	//运行的时候不要在golang中运行，因为会找不到静态文件夹的位置，所以我们要在命令行里运行
	//go build -o router_static && ./router_static  		-->-o 表示指定输出文件
	//启动之后我们进行一下测试，curl "http://127.0.0.1:8080/assets/a.html"
	//同理，完成下面两种
	r.Run()
}
