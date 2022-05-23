package main

import (
	"fmt"
	"reflect"
)

type neibu struct {
	a []int
}

type hehe struct {
	neibu
}

func main() {
	/*var a int = 1024

	valueOfA := reflect.ValueOf(a)

	var getA int = valueOfA.Interface().(int)

	var getA2 int = int(valueOfA.Int())

	fmt.Println(getA, getA2)*/

	var a = hehe{
		neibu: neibu{
			a: []int{10, 9, 8},
		},
	}

	valueOfA := reflect.ValueOf(a)

	fmt.Println(valueOfA.FieldByIndex([]int{0, 0})) // 只能打开嵌套结构,不能再继续访问 slice / map 这些内置类型

}
