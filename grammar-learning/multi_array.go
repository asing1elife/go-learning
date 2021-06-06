package main

import "fmt"

// 演示多维数组中的二维数组
func main() {
	// 声明二维数组
	var names [][]string

	// 准备两个一维数组
	firsts := []string{"王", "陈"}
	lasts := []string{"萌", "娴"}

	// 使用 append 函数，分别将两个一维数组添加到二维数组中
	names = append(names, firsts)
	names = append(names, lasts)

	fmt.Println(names)

	// 使用嵌套循环输出二维数组的内容
	for i := 0; i < len(names); i++ {
		for j := 0; j < len(names); j++ {
			fmt.Printf("names[%d][%d] = %s\n", i, j, names[i][j])
		}
	}
}
