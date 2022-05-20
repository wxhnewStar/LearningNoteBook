package main

import (
	"fmt"
)

type Dictionary struct {
	data map[interface{}]interface{} // 键值对都是 空接口 类型
}

// 根据键获取值
func (d *Dictionary) Get(key interface{}) (value interface{}) {
	return d.data[key]
}

func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

// 遍历所有的键值,如果回调返回值为 false, 停止遍历
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {

	if callback == nil {
		return
	}

	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

func (d *Dictionary) Clear() {
	// 前面讲过, 清空一个容器的办法就是让它重新 make 一个新容器
	d.data = make(map[interface{}]interface{})
}

func NewDictionary() *Dictionary {
	d := &Dictionary{}

	// 初始化 map
	d.Clear()
	return d
}

// 写一个函数用于判断一个 空接口 是什么类型
func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, " is int")
	case string:
		fmt.Println(v, " is string")
	case bool:
		fmt.Println(v, " is bool")
	default:
		fmt.Println(v, " is other")
	}
}

func main() {

	dict := NewDictionary()

	// 添加游戏数据
	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Don't Hungry", 24)

	// 获取值及打印值
	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite:", favorite)

	// 遍历所有的字典元素
	dict.Visit(func(k, v interface{}) bool {
		fmt.Println("visiting:")
		if v.(int) > 40 {
			fmt.Println(k, "is expensive")
			return true
		}
		fmt.Printf("%v is chep\n", k)
		return true
	})

	var val interface{} = 11
	printType(val)

	/*data := make(map[int]int)
	data[1] = 1
	data[2] = 2

	fmt.Println("开始测试")

	var dataRefer = data
	fmt.Println(dataRefer[1])

	fmt.Println("重新 make data 容器")

	data = make(map[int]int)
	data[3] = 3

	fmt.Println("After make, dataRefer——", dataRefer)
	fmt.Println("After make, dataRefer——", data)*/
}
