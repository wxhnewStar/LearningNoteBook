package main

import "fmt"

// Profile 人员档案
type Profile struct {
	Name    string
	Age     int
	Married bool
}

/* ------- 字符串转换哈希值 ------- */

func simpleHash(str string) (ret int) {
	// 将字符串转换为整数值
	for i := 0; i < len(str); i++ {
		// 取出字符
		c := str[i]

		// 将字符的 ASCII 相加
		ret += int(c)
	}
	return ret
}

/* --------- 查询键 ---------*/

// 查询键
type classicQueryKey struct {
	Name string
	age  int
}

// 计算查询键的哈希值
func (c *classicQueryKey) hash() int {
	return simpleHash(c.Name) + c.age
}

/* ------- 构建索引 ---------- */

// 创建哈希值到数据的索引关系
var mapper = make(map[int][]*Profile)

func buildIndex(list []*Profile) {
	// 遍历所有数据
	for _, profile := range list {

		// 构建数据的查询索引
		key := classicQueryKey{profile.Name, profile.Age}

		// 计算数据的哈希值,取出已经存在的记录
		existValue := mapper[key.hash()]

		// 将当前数据添加到已村子的记录切片中
		existValue = append(existValue, profile)

		// 将切片重新设置到映射中
		mapper[key.hash()] = existValue
	}
}

/* ------- 查询逻辑 ------- */
func queryData(name string, age int) {
	// 根据给定查询条件构建查询键
	keyToQuery := classicQueryKey{name, age}

	// 计算查询键的哈希值并查询,获得相同哈希值的所有结果集合
	resultList := mapper[keyToQuery.hash()]

	// 遍历结果集合
	for _, result := range resultList {
		if result.Name == name && result.Age == age {
			fmt.Println(result)
			return
		}
	}

	fmt.Println("Not found")
}

func main() {

	list := []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 21},
		{Name: "王麻子", Age: 21},
	}

	buildIndex(list)

	queryData("张三", 30)
}
