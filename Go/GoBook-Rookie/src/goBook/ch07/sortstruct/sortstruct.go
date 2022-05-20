package main

import (
	"fmt"
	"sort"
)

// 声明英雄的分类
type HeroKind int

// 定义 HeroKind 变量,类似于枚举
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

// 定义英雄名单的结构
type Hero struct {
	Name string
	Kind HeroKind
}

// 将英雄指针的切片定义为 Heros 类型
type Heros []*Hero

// 实现 sort.Interface 接口取元素数量方法
func (heros Heros) Len() int {
	return len(heros)
}

// 实现 sort.Interface 接口比较元素方法
func (heros Heros) Less(i, j int) bool {
	// 如果英雄的分类不同,优先对分类进行排序
	if heros[i].Kind != heros[j].Kind {
		return heros[i].Kind < heros[j].Kind
	}
	return heros[i].Name < heros[j].Name
}

func (heros Heros) Swap(i, j int) {
	heros[i], heros[j] = heros[j], heros[i]
}

func main() {
	// 准备英雄列表
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}

	//sort.Sort(heros)

	/*----- 使用 sort.Slice 类型 ------*/
	fmt.Println("using sort.Slice()")

	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})

	for _, v := range heros {
		fmt.Println("%+v", v)
	}

}
