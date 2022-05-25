package main

import "fmt"

// 定义一个结构体
type MyImplement struct {
}

// 实现 fmt.Stringer 的 String 方法
func (m *MyImplement) String() string {
	return "hi"
}

// 在函数中返回 fmt.Stringer 接口
func GetStringer() fmt.Stringer {
	// 赋 nil
	var s *MyImplement = nil

	return s
}

func main() {

	if GetStringer() == nil {
		fmt.Println("GetStringer() == nil")
	} else {
		fmt.Println("GetStringer() != nil")
	}
}
