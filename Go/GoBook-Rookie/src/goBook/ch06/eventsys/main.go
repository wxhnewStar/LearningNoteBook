package main

import (
	"fmt"
)

type Actor struct {
}

// OnEvent 为角色添加一个事件处理函数
func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

// GlobalEvent 全局事件
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func main() {
	a := new(Actor)

	RegisterEvent("OnSkill", a.OnEvent)

	RegisterEvent("OnSkill", GlobalEvent)

	CallEvent("OnSkill", 100)
}
