package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})

	//这里我们不使用run，使用自动化配置的拓展包
	//流程大致是：首先生成一个本地的密钥，然后把密钥发送给了一个证书颁发机构，获取一个私钥，之后本地要做一个私钥的验证，验证成功
	//会把私钥的信息储存起来，下次请求时使用私钥加密，达到一个自动下载证书的功能。
	autotls.Run(r, "www.itpp.tk")
}
