package search

func NumIslands(grid [][]byte) int {
	var (
		length = len(grid)
		width  = len(grid[0])
		count  = 0
	)
	if length == 0 {
		return 0
	}
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '1' {
				dfs(i, j, width, length, grid)
				count += 1
			}
		}
	}
	return count
}

func dfs(i, j, width, length int, grid [][]byte) {
	if i < 0 || i >= length || j < 0 || j >= width {
		return
	}
	for grid[i][j] == '1' {
		grid[i][j] = '0'
		dfs(i-1, j, width, length, grid) // up
		dfs(i+1, j, width, length, grid) // down
		dfs(i, j-1, width, length, grid) // left
		dfs(i, j+1, width, length, grid) // right
	}
}
