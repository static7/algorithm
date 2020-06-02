package power

import (
	"reflect"
	"testing"
)

// 求幂
// @author staitc7 <static7@qq.com>
func TestPowers(t *testing.T) {
	got := Powers(float64(2), 1)
	want := float64(2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("期待结果:%v, 实际结果:%v", want, got)
	}
}

// 基准测试
// @author staitc7 <static7@qq.com>
func BenchmarkPowers(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i <= b.N; i++ {
		Powers(2.0, 10)
	}
}
