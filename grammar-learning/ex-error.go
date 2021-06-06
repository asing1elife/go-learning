package main

import "fmt"

// 演示错误处理
func main() {
	result, msg := division(100, 10)
	fmt.Println(result, msg)

	result1, msg1 := division(100, 0)
	fmt.Println(result1, msg1)
}

// DivisionError 除法错误
type DivisionError struct {
	// 除数
	divisor int
	// 被除数
	beDivided int
}

// 实现内置的 error 接口中的 Error 函数
func (de *DivisionError) Error() string {
	return fmt.Sprintf("除数 %d, 被除数 %d, 被除数不能为 0", de.divisor, de.beDivided)
}

func division(divisor int, beDivided int) (result int, err string) {
	if beDivided == 0 {
		// 获取异常数据
		de := DivisionError{divisor, beDivided}

		return 0, de.Error()
	}

	return divisor / beDivided, ""
}
