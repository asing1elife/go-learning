package main

import "fmt"

// 演示切片，也就是动态（ 长度 ）数组，切片是对数组的抽象
func main() {
	// 初始化数组的简写形式其实就是创建了一个切片
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 输出完成切片
	printSlice(nums)
	// 输出索引 1 到 4 的元素
	printSlice(nums[1:4])
	// 输出索引 0 到 4 的元素，前置索引不指定时，默认为 0
	printSlice(nums[:4])
	// 输出索引 1 到数组长度的元素，后置索引不指定时，默认为 len(s)
	printSlice(nums[1:])

	// 内置函数 append() 可以向切片追加一个或多个元素
	nums = append(nums, 10, 11, 12)
	printSlice(nums)

	num1s := []int{13, 14, 15}
	// 内置函数 copy() 可以将新的切片按照索引的位置复制到原始切片中，原始切片中对应索引的元素会被替换
	copy(nums, num1s)
	printSlice(nums)
}

func printSlice(nums []int) {
	// 内置函数 len() 可以获取数组的长度
	// 内置函数 cap() 可以获取数组的容量
	fmt.Printf("数组的长度 -> %d , 容量 -> %d , 元素 -> %d \n", len(nums), cap(nums), nums)
}
