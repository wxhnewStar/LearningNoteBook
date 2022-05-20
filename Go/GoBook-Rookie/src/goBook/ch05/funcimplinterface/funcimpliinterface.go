package main

import "fmt"

type Invoker interface {
	Call(interface{})
}

// Struct 结构体类型
type Struct struct {
}

// Call 实现 Invoker 的 Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from Struct", p)
}

// FuncCaller 函数实现接口, 函数需要定义为类型
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	// 调用 f() 函数本体
	f(p)
}

func main() {
	// 声明接口变量
	var invoker Invoker

	s := new(Struct)
	invoker = s
	invoker.Call("hello")

	/* --------- 函数体调用接口 ----------- */

	// 将匿名函数转为 FuncCaller 类型,再赋值给接口
	invoker = FuncCaller(func(p interface{}) {
		fmt.Println("from function", p)
	})
	invoker.Call("hello2")
}
