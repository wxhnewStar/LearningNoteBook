package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getValue(filename, exceptSection, expectKey string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 使用reader读取文件
	reader := bufio.NewReader(file)
	// 当前读取的段的名字
	var sectionName string

	for {
		// 读取文件的一行
		lineStr, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// 切掉行左右两边的空白字符
		lineStr = strings.TrimSpace(lineStr)

		// 忽略空行
		if lineStr == "" {
			continue
		}

		// 忽略注释
		if lineStr[0] == ';' {
			continue
		}

		// 读取段和键值的代码
		if lineStr[0] == '[' && lineStr[len(lineStr)-1] == ']' {
			sectionName = lineStr[1 : len(lineStr)-1]
		} else if sectionName == exceptSection {
			pair := strings.Split(lineStr, "=")
			if len(pair) == 2 {
				key := strings.TrimSpace(pair[0])

				if key == expectKey {
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return ""
}

func main() {

	// "go run .\goBook\ch02\inireader\inireader.go .\goBook\CH02\inireader\example.ini core"

	args := os.Args

	if len(args) != 4 {
		println("usage: filename , desiredSection, desiredKey")
		log.Fatal("usage: filename , desiredSection, desiredKey")
	}
	var filename = args[1]
	var desiredSection = args[2]
	var desiredKey = args[3]

	val := getValue(filename, desiredSection, desiredKey)
	if val == "" {
		println("can not find this section & key")
	} else {
		println("the value is :", val)
	}
}
