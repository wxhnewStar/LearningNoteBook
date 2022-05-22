package main

import (
	"fmt"
	"time"
)

func main() {
	// 声明一个退出用的通道
	exit := make(chan int)

	// 打印开始
	fmt.Println("start")

	// 过 1 秒后,调用匿名函数
	time.AfterFunc(time.Second, func() {
		fmt.Println("one second has gone")
		exit <- 0
	})

	// 等待结束
	<-exit
}
