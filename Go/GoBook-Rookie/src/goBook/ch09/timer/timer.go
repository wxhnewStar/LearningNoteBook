package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个打点器,每 500 ms 触发一次
	ticker := time.NewTicker(time.Millisecond * 500)

	// 创建一个计时器, 2 s 后触发
	stopper := time.NewTimer(time.Second * 2)

	// 声明一个计数变量
	var i int

	// 不断检查通道情况
	for {
		// 多路复用通道
		select {
		case <-stopper.C:
			fmt.Println("stop")
			goto StopHere
		case <-ticker.C:
			i++
			fmt.Println("tick:", i)
		}
	}

	// 退出的标签,使用 goto 跳转
StopHere:
	fmt.Println("done")
}
