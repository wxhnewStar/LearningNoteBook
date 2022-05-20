package main

import (
	"fmt"
)

type ParseError struct {
	Filename string
	Line     int
}

// 重载了 Error() 方法,既可视为 error 接口?
func (e *ParseError) Error() string {
	return fmt.Sprintf("Filename: %s & Line: %d ", e.Filename, e.Line)
}

func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func main() {
	var e error

	e = newParseError("main.go", 1)

	fmt.Println(e.Error())

	switch detail := e.(type) {
	case *ParseError:
		fmt.Println(detail.Filename, detail.Line)
	default:
		fmt.Println("other error")
	}
}
