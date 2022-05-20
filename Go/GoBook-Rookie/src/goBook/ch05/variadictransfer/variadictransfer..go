package main

import (
	"fmt"
)

func rawPrint(rawList ...interface{}) {
	// 遍历可变参数切片
	for _, a := range rawList {
		fmt.Println(a)
	}
}

func print(slist ...interface{}) {
	rawPrint(slist...)
}

func main() {
	print(1, 2, 3)
}
