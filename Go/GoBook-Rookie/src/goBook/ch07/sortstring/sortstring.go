package main

import (
	"fmt"
	"sort"
)

type IntList []int

func (il *IntList) Swap(i, j int) {
	temp := (*il)[i]
	(*il)[i] = (*il)[j]
	(*il)[j] = temp
}

func (il *IntList) Len() int {
	return len(*il)
}

func (il *IntList) Less(i, j int) bool {
	return (*il)[i] < (*il)[j]
}

func main() {
	intList := IntList{5, 2, 4, 1, 3}

	fmt.Println("before :", intList)
	sort.Sort(&intList)
	fmt.Println("after :", intList)

	// 字符串切片 有 便捷排序可用
	// type StringSlice []string
	var strList = sort.StringSlice{"zlib", "hehe", "wocao"}

	fmt.Println("before :", strList)
	sort.Sort(strList)
	fmt.Println("after :", strList)

	// int 类型有同样的 IntSlice 类型
	// Sort.IntSlice
}
