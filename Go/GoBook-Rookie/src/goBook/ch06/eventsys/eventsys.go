package main

// 实例化一个通过字符串映射到函数切片的 map
var eventByName = make(map[string][]func(interface{}))

// 注册事件,提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	// 通过名字查找事件列表
	list := eventByName[name]

	// 在列表切片中添加函数
	list = append(list, callback)

	// 保存修改的事件列表切片
	eventByName[name] = list
}

// 调用事件
func CallEvent(name string, param interface{}) {
	list := eventByName[name]

	for _, callback := range list {
		// 传入参数调用回调函数
		callback(param)
	}
}
