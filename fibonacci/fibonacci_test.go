package fibonacci

import (
	"reflect"
	"testing"
)

// 测试
// @author staitc7 <static7@qq.com>
func TestFibonacci(t *testing.T) {
	got := Fibonacci(25)
	want := uint64(75025)
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("期待结果:%v, 实际结果:%v", want, got) // 测试失败输出错误提示
	}
}

// 基准测试
// @author staitc7 <static7@qq.com>
func BenchmarkFibonacci(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i <= b.N; i++ {
		Fibonacci(10)
	}
}
