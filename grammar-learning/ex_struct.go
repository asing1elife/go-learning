package main

import "fmt"

// 演示结构体
func main() {
	// 结构体可以按照内部变量的顺序进行初始化
	fmt.Println(Member{"asing1elife", "男", 27})
	// 也可以使用 key-value 的形式进行初始化
	fmt.Println(Member{name: "asing1elife", gender: "男", age: 27})
	// 在 key-value 形式中，对于没有指定的变量，默认值会是 0 或 空
	fmt.Println(Member{name: "asing1elife", gender: "男"})

	// 初始化结构体
	var member Member
	// 赋值
	member.name = "asing1elife"
	member.gender = "男"
	member.age = 27
	// 输出
	fmt.Printf("姓名：%s , 性别：%s , 年龄：%d \n", member.name, member.gender, member.age)
}

type Member struct {
	name   string
	gender string
	age    int
}
