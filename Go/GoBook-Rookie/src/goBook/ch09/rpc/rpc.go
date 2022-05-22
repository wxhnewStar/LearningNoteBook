package main

import (
	"errors"
	"fmt"
	"time"
)

/* --------- 客户端 -------- */

// 模拟 RPC 客户端的请求和接收信息封装
func RPCClient(ch chan string, req string) (string, error) {
	// 向服务器发送请求
	ch <- req

	// 等待服务器服务器返回
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("Time out error")
	}

}

/* -------- 服务端 --------- */

// 模拟 RPC 服务器端接收客户端请求和回应
func RPCServer(ch chan string) {
	for {
		// 接收客户端请求
		data := <-ch

		fmt.Println("Server received:", data)

		// 向客户端反馈已收到
		ch <- "roger"
	}
}

func main() {
	ch := make(chan string)

	go RPCServer(ch)

	recv, err := RPCClient(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Client received:", recv)
	}
}
