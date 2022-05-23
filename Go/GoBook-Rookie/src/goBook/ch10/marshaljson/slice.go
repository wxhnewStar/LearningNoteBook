package main

import (
	"bytes"
	"reflect"
)

func writeSlice(buff *bytes.Buffer, value reflect.Value) error {
	// 写入切片开始标记
	buff.WriteString("[")

	var err error

	for s := 0; s < value.Len(); s++ {
		sliceValue := value.Index(s)

		err = writeAny(buff, sliceValue)
		if err != nil {
			return err
		}

		if s < value.Len()-1 {
			buff.WriteString(",")
		}
	}

	buff.WriteString("]")

	return nil
}
