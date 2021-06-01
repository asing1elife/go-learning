package main

import (
	"fmt"
	"time"
)

// 演示并发
func main() {
	// 为什么只写一个 go say 没有任何输出
	go say("asing1elife")
	say("e1even")
}

func say(word string) {
	for i := 0; i < 5; i++ {
		// 睡一秒
		time.Sleep(1000 * time.Millisecond)

		fmt.Println(word)
	}
}
