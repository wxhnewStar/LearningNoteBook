package main

import (
	"fmt"
	"runtime"
)

type panicContext struct {
	function string
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()

		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error", err)
		}
	}()

	entry()
}

func main() {
	fmt.Println("运行前")

	ProtectRun(func() {
		fmt.Println("手动宕机前")

		// 使用 panic 传递上下文
		panic(&panicContext{
			"手动触发panic",
		})

		fmt.Println("手动宕机后")
	})

	// 故意造成空指针访问错误
	ProtectRun(func() {
		fmt.Println("赋值宕机前")

		var a *int
		*a = 1

		fmt.Println("赋值宕机后")
	})

	fmt.Println("运行后")
}
