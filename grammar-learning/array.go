package main

import "fmt"

// 演示数组
func main() {
	// 最简介的声明方式，虽然没有指定长度，但编译器会自动计算
	nums := []int{1, 2, 3}
	// 通过 ... 可以告诉编译器，该数组长度不确定
	num1s := [...]int{1, 2, 3}
	// 可以指定长度，如果赋值内容超出长度，在编译时就会报错
	num2s := [3]int{1, 2, 3}
	// 明确指定长度后，可以在初始化时为指定索引赋值
	num3s := [3]int{0: 1, 2: 3}

	fmt.Println("nums 数组长度", len(nums))
	fmt.Println("num1s 数组长度", len(num1s))
	fmt.Println("num2s 数组长度", len(num2s))
	fmt.Println("num3s 数组长度", len(num3s))

	for i := 0; i < len(num3s); i++ {
		// 因为 num3s 在初始化时没有指定索引 1 的值，所以该索引的值输出是 0
		fmt.Printf("num3s 当前索引 %d 的值是 %d\n", i, num3s[i])
	}
}
