package main

import (
	"fmt"
	"reflect"
)

func main() {
	// *int 的空指针
	var a *int
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())

	// nil 值
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())

	// *int 类型的空指针
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())

	// 实例化一个结构体
	s := struct {
	}{}

	// 尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid())

	// 尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(s).MethodByName("").IsValid())

	// 实例化一个 map
	m := map[int]int{}

	// 尝试从 map 中查找一个不存在的键
	fmt.Println("不存在的键:", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
}
