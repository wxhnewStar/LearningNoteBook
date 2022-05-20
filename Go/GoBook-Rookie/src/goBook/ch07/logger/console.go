package main

import (
	"fmt"
	"os"
)

// 命令行写入器
type consoleWriter struct {
}

// 实现 LogWriter 的 Write() 方法
func (c *consoleWriter) Write(data interface{}) error {
	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)

	_, err := os.Stdin.Write([]byte(str))

	return err
}

// 创建命令行写入器实例
func newConsolewriter() *consoleWriter {
	return &consoleWriter{}
}
