package main

import "fmt"

// 演示 range
func main() {
	nums := []int{1, 2, 3}

	sum := 0
	// 使用 range 操作数组/切片时可以分别获取到当前元素的索引和元素本身
	// 这里因为不需要使用索引，所以使用 _ 占位符来过滤掉
	for _, num := range nums {
		sum += num
	}

	fmt.Printf("sum -> %d \n", sum)

	kvs := map[string]string{"name": "asing1elife", "gender": "male"}
	// 使用 range 操作 map 可以获取键值对
	for key, value := range kvs {
		fmt.Printf("key -> %s, value -> %s \n", key, value)
	}
}
