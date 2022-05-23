package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int

	typeOfA := reflect.TypeOf(a)

	aIns := reflect.New(typeOfA)

	var aa = aIns.Elem()

	fmt.Println(aa.Type(), aa.Kind())

	fmt.Println(aIns.Type(), aIns.Kind())
}
