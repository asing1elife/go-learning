package main

import "github.com/gin-gonic/gin"

// 演示最基础的请求
func main() {
	// 实例化请求对象
	engine := gin.Default()

	// 定义请求
	engine.GET("/hello", func(context *gin.Context) {
		// 返回数据
		context.JSON(200, gin.H{
			"msg": "Hi asing1elife",
		})
	})

	// 指定端口并运行，冒号不可或缺
	engine.Run(":8081")
}
