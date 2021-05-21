package main

import "fmt"

// 演示接口
func main() {
	var human Human

	human = new(Student)
	human.selfIntroduction()

	human = new(Teacher)
	human.selfIntroduction()
}

// Human 接口中不能直接定义变量
type Human interface {
	selfIntroduction()
}

type Student struct {
}

type Teacher struct {
}

func (student Student) selfIntroduction() {
	fmt.Println("我是学生")
}

func (teacher Teacher) selfIntroduction() {
	fmt.Println("我是老师")
}
