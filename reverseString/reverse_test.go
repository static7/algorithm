package reverseString

import (
	"reflect"
	"testing"
)

// 反转字符串测试
// @author staitc7 <static7@qq.com>
func TestReverse(t *testing.T) {
	got := ReverseString([]byte("hello"))
	want := []byte("olleh")
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}

}

//基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		ReverseString([]byte("hello"))
	}
}
