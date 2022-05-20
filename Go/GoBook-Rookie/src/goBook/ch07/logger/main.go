package main

import "fmt"

// 创建日志器
func createLogger() *Logger {

	l := NewLogger()

	cw := newConsolewriter()

	l.RegisterWriter(cw)

	fw := newFileWriter()

	// 设置文件名
	if err := fw.SetFile("./log.log"); err != nil {
		fmt.Println(err)
	}

	// 注册文件写入器到日志器中
	l.RegisterWriter(fw)

	return l
}

func main() {
	l := createLogger()

	l.Log("hello")
}
