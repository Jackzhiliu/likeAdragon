package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个 Gin 实例
	r := gin.Default()

	// 定义一个简单的 GET 路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 运行服务器，监听 8080 端口
	r.Run(":8080")
}