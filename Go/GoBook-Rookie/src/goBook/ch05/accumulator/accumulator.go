package main

import "fmt"

func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value++

		// 返回一个累加值
		return value
	}
}

func main() {
	//  创建一个累加器,初始值为 1
	accumulator := Accumulate(1)

	// 累加 1 并打印
	fmt.Println(accumulator())

	fmt.Println(accumulator())

	// 创建一个累加其,初始值为10
	accumulator2 := Accumulate(10)

	// 累加 1 并打印
	fmt.Println(accumulator2())
}
