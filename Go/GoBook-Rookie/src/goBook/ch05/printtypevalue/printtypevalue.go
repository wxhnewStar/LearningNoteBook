package main

import (
	"bytes"
	"fmt"
)

func printTypeValue(slist ...interface{}) string {

	var b bytes.Buffer

	for _, s := range slist {
		// 将 interface{} 类型格式化为字符串
		str := fmt.Sprintf("%v", s)

		var typeString string

		// 对 s 进行类型断言
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		case float64:
			typeString = "float64"
		default:
			typeString = "unpointer"
		}

		b.WriteString("value: ")

		b.WriteString(str)

		b.WriteString(" type: ")

		b.WriteString(typeString)

		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	fmt.Println(printTypeValue("hehe", 1000, float64(10.02), []int{1, 2, 3}))
}
