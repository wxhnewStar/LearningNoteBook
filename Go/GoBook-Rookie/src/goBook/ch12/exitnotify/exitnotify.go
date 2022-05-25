package main

import (
	"fmt"
	"net"
	"time"
)

func socketRecv(conn net.Conn, exitChan chan string) {
	// 创建一个接收缓冲
	buff := make([]byte, 1024)

	var count int
	for {
		// 从套接字中读取数据
		_, err := conn.Read(buff)

		// 需要结束接收,退出循环
		if err != nil {
			fmt.Println(err)
			break
		}
		count++
		fmt.Printf("Read Count %d :%s\n", count, string(buff))
	}

	exitChan <- "recv exit"
}

func main() {
	// 连接一个地址
	conn, err := net.Dial("tcp", "www.baidu.com:80")

	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建退出通道
	exit := make(chan string)

	// 并发执行套接字接收
	go socketRecv(conn, exit)

	// 在接收时,等待 1 秒
	time.Sleep(time.Second * 10)

	// 主动关闭套接字
	defer conn.Close()

	// 等待 goroutine 退出完毕
	fmt.Println(<-exit)
}
