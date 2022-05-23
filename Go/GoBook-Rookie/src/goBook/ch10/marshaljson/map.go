package main

import (
	"bytes"
	"reflect"
)

func writeMap(buff *bytes.Buffer, value reflect.Value) error {
	buff.WriteString("[")

	var err error

	iter := value.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		buff.WriteString("\"")

		err = writeAny(buff, k)
		if err != nil {
			return err
		}

		buff.WriteString("\":\"")

		err = writeAny(buff, v)
		if err != nil {
			return err
		}

		buff.WriteString("\"")

		buff.WriteString(",")
	}
	buff.WriteString("]")
	return nil
}
