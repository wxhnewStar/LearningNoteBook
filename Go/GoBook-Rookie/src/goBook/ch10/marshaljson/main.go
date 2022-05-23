package main

import "fmt"

// 声明技能结构
type Skill struct {
	Name  string
	Level int
}

// 声明角色结构
type Actor struct {
	Name  string
	Age   int
	myMap map[int]int

	Skills []Skill
}

func main() {
	// 填充基本角色数据
	a := Actor{
		Name: "cow boy",
		Age:  37,
		Skills: []Skill{
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash your dog eye", Level: 2},
			{Name: "Time to have lunch", Level: 3},
		},
		myMap: map[int]int{
			10: 100,
			20: 200,
		},
	}

	if result, err := MarShalJson(a); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
