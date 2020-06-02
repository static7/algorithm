package pascal_triangle

import (
	"reflect"
	"testing"
)

// 杨辉三角测试
// @author staitc7 <static7@qq.com>
func TestGenerate(t *testing.T) {
	got := Generate(5)
	want := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("期待结果:%v, 实际结果:%v", want, got) // 测试失败输出错误提示
	}
}

// 基准测试
// @author staitc7 <static7@qq.com>
func BenchmarkGenerate(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Generate(5)
	}
}

// 杨辉三角测试
// @author staitc7 <static7@qq.com>
func TestGenerate2(t *testing.T) {
	got := Generate2(3)
	want := []int{1, 3, 3, 1}
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("期待结果:%v, 实际结果:%v", want, got) // 测试失败输出错误提示
	}
}
