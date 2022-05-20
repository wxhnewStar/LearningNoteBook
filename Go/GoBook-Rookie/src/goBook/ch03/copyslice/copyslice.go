package main

import "fmt"

func main() {
	const elementCount = 1000

	srcData := make([]int, elementCount)

	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	// 这里创建的是一个 引用
	refData := srcData

	copyData := make([]int, elementCount)
	copy(copyData, srcData)

	// 修改原始数据的第一个元素
	srcData[0] = 999

	// 打印引用切片的第一个元素
	fmt.Println(refData[0])

	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])

	// 每次都是从 copyData 开始位置开始复制
	copy(copyData, srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
	fmt.Println("")
}
