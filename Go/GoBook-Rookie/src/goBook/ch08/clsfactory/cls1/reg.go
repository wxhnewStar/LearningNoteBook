package cls1

import (
	"fmt"
	"goBook/ch08/clsfactory/base"
)

type Class1 struct {
}

func (c *Class1) Do() {
	fmt.Println("Class1")
}

func init() {

	// 在启动时注册 class1 工厂
	base.Register("Class1", func() base.Class {
		return new(Class1)
	})
}
