package fibonacci

// 斐波那契数列
// @author staitc7 <static7@qq.com>
func Fibonacci(n uint64) uint64 {
	if n < 2 {
		return n
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}

}
