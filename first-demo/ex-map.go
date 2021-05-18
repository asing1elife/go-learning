package main

import "fmt"

// 演示 map
func main() {
	// 定义 map， int 是 key 的类型，string 是 value 的类型
	var kvs = map[int]string{}

	// 初始化
	kvs[0] = "asing1elife"
	kvs[1] = "e1even"

	show(kvs)

	// 删除指定 key
	delete(kvs, 0)

	fmt.Println("删除某个元素后")
	show(kvs)
}

func show(kvs map[int]string) {
	for key, value := range kvs {
		fmt.Printf("key -> %d, value -> %s \n", key, value)
	}
}
