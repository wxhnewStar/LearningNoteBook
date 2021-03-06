package main

import (
	"bytes"
	"reflect"
)

func MarShalJson(v interface{}) (string, error) {
	// 准备一个缓冲
	var b bytes.Buffer

	// 将任意值转换为 JSON 并输出到缓冲中
	if err := writeAny(&b, reflect.ValueOf(v)); err == nil {
		return b.String(), nil
	} else {
		return "", err
	}
}
