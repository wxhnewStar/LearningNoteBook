package main

import (
	"fmt"
	"strings"
)

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

//  字符串处理函数,传入字符串切片和处理链, 切片传的是引用
func StringProccess(list []string, chain []func(string) string) {
	// 遍历每一个字符串
	for index, str := range list {
		// 第一个需要处理的字符串
		result := str

		// 遍历每一个处理链
		for _, proc := range chain {
			// 输入一个字符串进行处理,返回数据作为下一个处理链的输入
			result = proc(result)
		}

		// 将结果放回切片
		list[index] = result
	}
}

func main() {
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
	}

	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	// 处理字符串
	StringProccess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}
}
