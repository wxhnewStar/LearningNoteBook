package main

import (
	"bytes"
	"reflect"
)

func writeStruct(buff *bytes.Buffer, value reflect.Value) error {
	// 取值的类型对象
	valueType := value.Type()

	// 写入结构体左大括号
	buff.WriteString("{")

	var err error

	// 遍历结构体的所有值
	for i := 0; i < value.NumField(); i++ {
		// 获取每个字段的字段值
		fieldValue := value.Field(i)

		fieldType := valueType.Field(i)

		// 写入字段左双引号
		buff.WriteString("\"")

		// 写入字段值
		buff.WriteString(fieldType.Name)

		// 写入字段名右双引号和冒号
		buff.WriteString("\":")

		// 写入字段值
		err = writeAny(buff, fieldValue)
		if err != nil {
			return err
		}

		//每个字段尾部写入逗号,最后一个不添加
		if i < value.NumField()-1 {
			buff.WriteString(",")
		}
	}

	// 写入结构体右大括号
	buff.WriteString("}")

	return nil
}
