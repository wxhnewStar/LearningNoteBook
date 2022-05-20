package main

import (
	"fmt"
)

type Player struct {
	name        string
	HealthPoint int
	MagicPoint  int
}

func main() {
	var p1 = Player{
		name:        "xxxx",
		HealthPoint: 0,
		MagicPoint:  0,
	}

	fmt.Println(p1.name)

	p2 := &Player{}

	fmt.Println(p2.name)
}
