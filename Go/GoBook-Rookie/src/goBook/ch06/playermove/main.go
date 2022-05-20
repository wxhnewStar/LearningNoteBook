package main

import "fmt"

func main() {
	// 实例化玩家对象,并设速度为 0.5
	p := NewPlayer(0.5)

	// 让玩家移动到 3,1
	p.MoveTo(Vec2{3, 1})

	for !p.IsArrived() {
		p.Update()

		fmt.Println(p.Pos())
	}
}
