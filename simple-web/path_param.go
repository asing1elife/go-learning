package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 演示携带路径参数的请求
func main() {
	engine := gin.Default()

	// :name 表示定义一个名称为 name 的路径参数
	// 这种定义方式存在一个弊端：如果不传值，则无法响应该请求
	engine.GET("/hello/:name", func(context *gin.Context) {
		// 获取参数
		name := context.Param("name")

		context.JSON(http.StatusOK, gin.H{
			"msg": "Hi " + name,
		})
	})

	// *name 也是用于定义路径参数，这种方式及时该位置不传值，也可以响应到该请求
	// 该请求格式可以响应以下三种请求：
	// 1. /bye/asing1elife
	// 2. /bye/
	// 3. /bye TIP: 除非本身没有定义 /bye 的请求，那么 /bye 请求会自动重定向到 /bye/
	engine.GET("/bye/*name", func(context *gin.Context) {
		name := context.Param("name")

		context.JSON(http.StatusOK, gin.H{
			"msg": "Bye " + name,
		})
	})

	engine.Run(":8081")
}
