package fibonacci

import (
	"reflect"
	"testing"
)

// 爬楼梯
// @author staitc7 <static7@qq.com>
func TestClimbStairs(t *testing.T) {
	got := ClimbStairs(4)
	want := 5
	if !reflect.DeepEqual(want, got) {
		t.Errorf("期待结果:%v, 实际结果:%v", want, got)
	}
}

// 基准测试
// @author staitc7 <static7@qq.com>
func BenchmarkClimbStairs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i <= b.N; i++ {
		ClimbStairs(5)
	}
}
