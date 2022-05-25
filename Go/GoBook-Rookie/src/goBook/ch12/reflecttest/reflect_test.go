package reflecttest

import (
	"reflect"
	"testing"
)

// 声明一个结构体,拥有一个字段
type data struct {
	Hp int
}

func BenchmarkNativeAssign(b *testing.B) {
	// 实例化结构体
	v := data{Hp: 2}

	// 停止基准测试的计时器
	b.StopTimer()
	// 重置基准测试的计时器
	b.ResetTimer()

	// 重新启动基准测试计时器
	b.StartTimer()

	// 根据基准测试数据进行循环测试
	for i := 0; i < b.N; i++ {
		// 结构体成员赋值测试
		v.Hp = 3
	}
}

func BenchmarkReflectAssign(b *testing.B) {
	// 实例化结构体
	v := data{Hp: 2}

	// 取出结构体指针的反射值对象并取其元素
	vv := reflect.ValueOf(&v).Elem()

	f := vv.FieldByName("Hp")

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		f.SetInt(3)
	}
}

/* ----- test 结构体成员赋值 ------- */
func BenchmarkReflectfindFieldAndAssign(b *testing.B) {
	v := data{Hp: 2}

	vv := reflect.ValueOf(&v).Elem()

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		vv.FieldByName("Hp").SetInt(3)
	}
}

/* ----- 调用函数对比 ----- */

func foo(v int) {

}

func BenchmarkNativeCall(b *testing.B) {

	for i := 0; i < b.N; i++ {
		foo(0)
	}
}

func BenchmarkReflectCall(b *testing.B) {

	v := reflect.ValueOf(foo)

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		v.Call([]reflect.Value{reflect.ValueOf(0)})
	}
}
