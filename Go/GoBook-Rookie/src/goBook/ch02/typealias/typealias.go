package main

import (
	"fmt"
	"reflect"
)

type Brand struct {
}

func (t Brand) Show() {

}

type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

func main() {
	var a Vehicle

	a.FakeBrand.Show()

	// a.Show()  这一行会报错,因为两个类型都有 Show() 方法,产生了歧义,证明 FakeBrand 的本质确实是 Brand 类型

	ta := reflect.TypeOf(a)

	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)

		fmt.Printf("FieldName: %v , FieldType: %v\n", f.Name, f.Type.Name())
	}
}