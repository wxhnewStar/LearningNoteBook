package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	// 设置图片大小
	const size = 300

	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))

	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	for x := 0; x < size; x++ {
		// 让 sin 的值范围在 0 ~ 2Pi 之间
		s := float64(x) / size * 2 * math.Pi

		// sin 的幅度为一半的像素, 向下偏移一半像素并反转
		y := size/2 - math.Sin(s)*size/2

		// 用黑色绘制 sin 轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}

	// 创建文件,把图像输出为文件
	file, err := os.Create("sin.png")

	if err != nil {
		log.Fatal(err)
	}

	// 使用 png 格式将数据写入文件
	png.Encode(file, pic)

	file.Close()
}
