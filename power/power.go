package power

// 求幂
// @author staitc7 <static7@qq.com>
func Powers(x float64, n int) float64 {
	switch n {
	case 0:
		return 1
	case 1:
		return x
	}
	if n < 0 {
		return Powers(1/x, -n)
	}
	y := Powers(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
