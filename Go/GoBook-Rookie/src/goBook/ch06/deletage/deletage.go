package main

import "fmt"

type class struct {
}

func (c *class) Do(v int) {
	fmt.Println("Call method do:", v)
}

func funcDo(v int) {
	fmt.Println("Call function do:", v)
}

func main() {
	var delegate func(int)

	c := new(class)

	delegate = c.Do

	delegate(10)

	delegate = funcDo

	delegate(10)
}
