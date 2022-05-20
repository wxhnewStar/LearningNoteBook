package main

import "fmt"

// 创建一个玩家生成器,输入名称,输出生成器
func playerGen(name string) func(mode string) (string, string, int) {

	// 血量一直为 150
	hp := 150

	// 返回创建的闭包
	return func(mode string) (string, string, int) {
		return mode, name, hp
	}
}

func main() {
	// 创建一个玩家生成器
	generator := playerGen("high noon")

	// 返回玩家的名字和血量
	mode, name, hp := generator("warrior")

	// 打印值
	fmt.Println(mode, name, hp)
}
