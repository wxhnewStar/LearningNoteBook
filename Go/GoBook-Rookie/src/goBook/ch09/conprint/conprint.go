package main

import "fmt"

func printer(c chan int) {
	// 开始无限循环等待数据
	for {
		data := <-c

		if data == 0 {
			break
		}

		fmt.Println(data)
	}

	// 通知 main 已经结束循环了
	c <- 0
}

func main() {
	c := make(chan int)

	go printer(c)

	for i := 1; i <= 10; i++ {
		c <- i
	}

	// 通知并发的 printer 结束循环
	c <- 0

	// 等待 printer 结束
	<-c
}
