package main

import "fmt"

// 演示通道
func main() {
	array := []int{1, 2, 3, 4, 5}

	// 需要先建立一个通道
	channel := make(chan int)
	// [:len(array)/2] 表示索引从 0 到数组长度除以 2 的位置，也就是 0 - 2 ，但不包含 2 ，具体元素覆盖为 1 、2
	go sum(array[:len(array)/2], channel)
	// [len(array)/2:] 表示索引从数组长度除以 2 的位置，也就是 2 一直到数组末尾，也就是 2 - 4 ，具体元素覆盖为 3 、4 、5
	go sum(array[len(array)/2:], channel)

	// 将通道中的值取出
	sum1, sum2 := <-channel, <-channel

	fmt.Println(sum1, sum2)

	fmt.Println("缓冲区演示 ->")

	// 建立一个通道，缓冲区设置为 2 ，表示只能存两个值
	buffer := make(chan int, 2)

	// 如果不设置缓冲区，下面这种赋值方式，会直接报错
	// 因为不存在缓冲区的通道，在赋值后必须立即有接收方来接收值
	buffer <- 1
	buffer <- 2

	sum3, sum4 := <-buffer, <-buffer

	fmt.Println(sum3, sum4)

	fmt.Println("通道关闭演示 ->")

	channel1 := make(chan int)

	go closeChannel(array, channel1)

	// range 函数可以用于遍历通道，不过不同于数组，返回的数据不存在索引
	for c := range channel1 {
		fmt.Println(c)
	}
}

// channel 是一个通道，用于传递并发线程中的数据
func sum(array []int, channel chan int) {
	sum := 0

	// 使用 range 来获取数组元素，_ 表示占位符，用于跳过索引
	for _, item := range array {
		sum += item
	}

	// 表示将 sum 的值传递到通道中
	channel <- sum
}

func closeChannel(array []int, channel chan int) {
	for _, item := range array {
		if item == 3 {
			// close 函数可以用于关闭通道
			close(channel)
			// 通道关闭后，必须退出循环，否则会报错
			break
		}

		channel <- item
	}
}
