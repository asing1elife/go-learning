package main

import "fmt"

// 在 Go 中，函数和方法是两个不同的概念
func main() {
	// 定义局部变量时，如果不指定默认值，则必须使用 var 进行声明
	var user User
	user.name = "asing1elife"

	fmt.Println(user.getName())
}

// User 结构体，也就是 Java 中的实体类
type User struct {
	name string
}

// 这个函数其实就是 User 结构体的方法，也就是 Java 中实体类的 Getter / Setter 方法的体现
func (user User) getName() string {
	return user.name
}
