package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float32
	ResX, ResY int
}

type Battery struct {
	Capacity int
}

func genJsonData() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery{
			2910,
		},
		true,
	}

	jsonData, _ := json.Marshal(raw)

	return jsonData
}

func main() {
	jsonData := genJsonData()

	fmt.Println(string(jsonData))

	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	// 反序列化到 screenAndTouch
	json.Unmarshal(jsonData, &screenAndTouch)

	// 输出 screenAndTouch 的详细结构
	fmt.Printf("%+v\n", screenAndTouch)

	// 只需要电池和指纹识别信息的结构和实例
	batteryAndTouch := struct {
		Battery
		HasTouchID bool // 发现只要是字母一样即可,不用管大小写
	}{}

	json.Unmarshal(jsonData, &batteryAndTouch)

	fmt.Printf("%+v\n", batteryAndTouch)
}
