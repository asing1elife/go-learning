package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("for i :", i)
	}

	i := 0
	// for 循环可以直接简写成 while 格式
	for i <= 10 {
		fmt.Println("while i :", i)
		i++
	}

	// 声明数组的方式
	arrs := []string{"tom", "jerry"}
	// 通过 range 可以遍历数据的索引和内容
	for i, arr := range arrs {
		fmt.Println(i, arr)
	}
}
