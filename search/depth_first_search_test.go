package search

import (
	"reflect"
	"testing"
)

// 测试深度优先搜索
// @author staitc7 <static7@qq.com>
func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	got := NumIslands(grid)
	want := 1
	if !reflect.DeepEqual(want, got) {
		t.Errorf("期待结果:%v, 实际结果:%v", want, got)
	}
}

// 基线测试
// @author staitc7 <static7@qq.com>
func BenchmarkNumIslands(b *testing.B) {
	b.ReportAllocs()
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	for i := 0; i <= b.N; i++ {
		NumIslands(grid)
	}
}
