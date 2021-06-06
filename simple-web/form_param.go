package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 演示携带表单参数的请求
func main() {
	engine := gin.Default()

	engine.POST("/user", func(context *gin.Context) {
		// 可以直接获取到表单提交中对应的字段值
		name := context.PostForm("name")
		// 获取参数的同时，判断参数是否传入
		gender, hasGender := context.GetPostForm("gender")

		if !hasGender {
			gender = "male"
		}

		context.JSON(http.StatusOK, gin.H{
			"name":   name,
			"gender": gender,
		})
	})

	engine.POST("/article", func(context *gin.Context) {
		// 如果要获取 RequestBody 中的数据，则需要提前定义结构体
		var article Article
		// 再通过 BindJSON() 进行绑定
		err := context.BindJSON(&article)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(article)

		context.JSON(http.StatusOK, article)
	})

	engine.Run(":8081")
}

type Article struct {
	// 在结构体中，如果想要成功接收到 JSON 格式的参数，首先变量名称需要大写，其次还需要在变量后标明希望接收的字段格式和 key
	Title  string `json:"title"`
	Author string `json:"author"`
}
