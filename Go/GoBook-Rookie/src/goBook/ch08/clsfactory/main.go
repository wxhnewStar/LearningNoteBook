package main

import (
	"goBook/ch08/clsfactory/base"
	_ "goBook/ch08/clsfactory/cls1"
	_ "goBook/ch08/clsfactory/cls2"
)

func main() {

	// 根据字符串动态创建一个 Class1 实例
	c1 := base.Create("Class1")
	c1.Do()

	c2 := base.Create("Class2")
	c2.Do()
}
