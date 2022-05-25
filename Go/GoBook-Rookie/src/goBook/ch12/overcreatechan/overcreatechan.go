package main

import (
	"fmt"
	"runtime"
)

// 一段耗时的计算函数
func consumer(ch chan int) {
	// 无限获取数据的循环
	for {
		data := <-ch
		fmt.Println(data)

		if data == 0 {
			break
		}
	}
}

func main() {
	// 创建一个传递数据用的通道
	ch := make(chan int)

	for {
		// 空变量,啥也不用做
		var dummy string

		// 获取输入,模拟进程持续运行
		fmt.Scan(&dummy)

		if dummy == "quit" {
			for i := 0; i < runtime.NumGoroutine()-1; i++ {
				ch <- 0
			}
			continue
		}

		go consumer(ch)

		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}
