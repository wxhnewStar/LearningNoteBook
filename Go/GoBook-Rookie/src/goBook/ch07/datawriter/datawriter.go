package main

import "fmt"

type Datawriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构,用于实现 DataWriter
type file struct {
}

// WriteData 实现 DataWriter 接口的 WriteData() 方法
func (f *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func main() {
	var f = new(file)

	var dataWriteInt Datawriter

	dataWriteInt = f

	dataWriteInt.WriteData("data")
}
