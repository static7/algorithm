package fibonacci

// 爬楼梯
// @author staitc7 <static7@qq.com>
func ClimbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	one, two := 1, 2
	for i := 3; i <= n; i++ {
		one, two = two, one+two
	}

	return two
}
