package main

import "fmt"

// 具备刷脸特性的接口
type ContainUseFaceID interface {
	CanUseFaceID()
}

// 具备被偷特性的接口
type ContainStolen interface {
	Stolen()
}

type Alipay struct {
}

// CanUseFaceID 为 Alipay 添加 CanUseFaceID 方法,表示电子支付方法支持刷脸
func (a *Alipay) CanUseFaceID() {

}

// Cash 现金支付方式
type Cash struct {
}

// 为 Cash 添加 Stolen() 方法,表示现金支付方式会出现偷窃情况
func (c *Cash) Stolen() {

}

// 打印支付方式具有的特点
func print(payMethod interface{}) {

	switch payMethod.(type) {
	case ContainStolen:
		fmt.Printf("%T can use faceid\n", payMethod)
	case ContainUseFaceID:
		fmt.Printf("%T may be stolen\n", payMethod)
	}
}

func main() {
	print(new(Alipay))

	print(new(Cash))
}
