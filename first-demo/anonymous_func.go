package main

import (
	"fmt"
	"math"
)

// 匿名函数
func main() {
	// 该函数体在 main() 函数中被定义
	result := func(num float64) float64 {
		return math.Abs(num)
	}

	// result() 不是变量，而是函数实参
	fmt.Println("当前值的绝对值是", result(-3))
}
