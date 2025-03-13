package main

import (
	"fmt"
	"log"
	"likeadragon/database"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	database.InitDatabase()

	// 测试连接
	sqlDB, err := database.DB.DB() // 获取原始数据库连接对象
	if err != nil {
		log.Fatal("Error getting database object:", err)
	}

	// 测试数据库连接
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	} else {
		fmt.Println("数据库连接成功")
	}

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
