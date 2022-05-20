package main

import (
	"fmt"
	"sync"
)

func main() {
	// 定义的方式就是直接声明就行,不需要 make
	var scene sync.Map

	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	fmt.Println(scene.Load("hehe")) // 找不到返回 nil,flase ,找得到返回 value,true

	scene.Delete("london")

	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
