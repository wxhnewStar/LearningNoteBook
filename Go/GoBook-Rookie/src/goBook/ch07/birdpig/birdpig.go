package main

import (
	"fmt"
)

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct {
}

func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

func (b *bird) Walk() {
	fmt.Println("birdL: walk")
}

type pig struct {
}

func (p *pig) Walk() {
	fmt.Println("pig: walk")
}
func main() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}

	for name, obj := range animals {
		// 判断对象是否为飞行物
		f, isFlyer := obj.(Flyer)
		// 判断对象是否为行走动物
		w, isWalker := obj.(Walker)

		fmt.Printf("name: %s, isFlyer:%v isWalker:%v\n", name, isFlyer, isWalker)

		if isFlyer {
			f.Fly()
		}

		if isWalker {
			w.Walk()
		}
	}
}
