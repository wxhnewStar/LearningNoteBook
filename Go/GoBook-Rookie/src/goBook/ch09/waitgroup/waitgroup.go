package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// 声明一个等待组
	var wg sync.WaitGroup

	// 准备一系列的网站地址
	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https;//www.golangtc.com/",
	}

	// 遍历这些地址
	for _, url := range urls {
		// 每一个任务开始时,将等待组增加 1
		wg.Add(1)

		// 开启一个并发
		go func(url string) {

			// 使用 defer,表示函数完成时将等待组减 1
			defer wg.Done()

			// 使用给定的 url 对其进行 http 访问
			_, err := http.Get(url)

			// 访问完成后,打印地址和可能发生的错误
			fmt.Println("Get:"+url, " ,err: ", err)

			time.Sleep(time.Second * 2)
		}(url)
	}

	// 等待所有的任务完成
	wg.Wait()

	fmt.Println("all get done")
}
