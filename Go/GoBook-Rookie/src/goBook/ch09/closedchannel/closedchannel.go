package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 0
	ch <- 1

	close(ch)

	// 关闭后依然可以访问内部数据
	for i := 0; i < cap(ch)+1; i++ {
		v, ok := <-ch

		fmt.Println(v, ok)
	}
}
