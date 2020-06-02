package pascal_triangle

// 杨辉三角
// @author staitc7 <static7@qq.com>
func Generate(numRows int) [][]int {
	arr := make([][]int, numRows)
	if numRows == 0 {
		return nil
	}
	for i := 0; i < numRows; i++ {
		arr[i] = make([]int, i+1)
		arr[i][0] = 1
		for j := 1; j <= i; j++ {
			if j == i {
				arr[i][j] = 1
				continue
			}
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
		}
	}
	return arr
}

// 杨辉三角2 笨方法
// @author staitc7 <static7@qq.com>
func Generate2(rowIndex int) []int {
	numRows := rowIndex + 1
	arr := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]int, i+1)
		arr[i][0] = 1
		for j := 1; j <= i; j++ {
			if j == i {
				arr[i][j] = 1
				continue
			}
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
		}
	}
	return arr[rowIndex]
}
