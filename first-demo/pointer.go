package main

import "fmt"

// 演示指针
func main() {
	// 定义普通变量
	var num int = 1
	// 定义指针变量
	var numPointer *int

	// 使用 & 可以获取到普通变量的地址，将其指向指针变量
	numPointer = &num

	fmt.Printf("num 的值：%d, 地址：%d \n", num, &num)
	// 要获取指针变量的地址而不是值，使用 * 即可
	fmt.Printf("numPointer 的值：%d, 地址：%d \n", *numPointer, numPointer)

}
