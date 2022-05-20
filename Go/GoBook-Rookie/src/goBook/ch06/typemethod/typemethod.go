package main

import "fmt"

type myInt int

func (m myInt) IsZero() bool {
	return m == 0
}

func (m myInt) Add(other int) int {
	return other + int(m)
}

func main() {
	var b myInt

	fmt.Println(b.IsZero())

	b = 1

	fmt.Println(b.Add(2))
}
