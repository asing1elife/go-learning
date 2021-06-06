package main

import "fmt"

// 演示函数调用
func main() {
	num1, num2 := 1, 2

	res := judgeMax(num1, num2)
	min, max := sortNum(num1, num2)

	fmt.Println("当前最大值是", res)
	fmt.Println("当前值的排列顺序是", min, max)
}

// judgeMax 函数形参必须指定类型
// 单个返回值的类型指定可以省略括号
func judgeMax(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

// sortNum 对传入参数进行排序
func sortNum(num1, num2 int) (int, int) {
	if num1 > num2 {
		return num2, num1
	}

	return num1, num2
}
