package main

import "fmt"

func main() {
	// _ 是内置的只读变量，用于排除多余的返回值
	_, a, b := getNumbers()

	fmt.Println(a, b)
}

/**
 函数存在返回值时，需要在函数体后使用 (varType ...) 的方式指定返回值的类型，否则会报错
 */
func getNumbers() (int, int, int) {
	// 形同 var a, b, c = 1, 2, 3
	a, b, c := 1, 2, 3

	return a, b, c
}
