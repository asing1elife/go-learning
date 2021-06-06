package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 演示携带请求参数的请求
func main() {
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		// 如果请求格式形如： /hello?name=xxx&gender=xxx&age=xxx ，以下两种方式都可以获取到参数值
		// DefaultQuery() 表示如果请求中没有传入 name 参数，那么将使用设定好的默认值
		name := context.DefaultQuery("name", "asing1elife")
		// Query() 表示直接尝试获取参数，如果没有传入，将为空字符串
		gender := context.Query("gender")
		// getQuery 在获取参数时，还会返回请求中是否传入该参数的标记
		age, hasAge := context.GetQuery("age")

		// 没有传入年龄参数
		if !hasAge {
			age = "18"
		}

		// String() 表示直接返回字符串
		context.String(http.StatusOK, "Hi %s , you gender is %s , age is %s", name, gender, age)
	})

	engine.Run(":8081")
}
