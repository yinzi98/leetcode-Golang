给定一个只含非负整数的 m x n 网格，找到一条从左上角到右下角的可以使数字之和最小的路径。

注意: 每次只能向下或者向右移动一步。

示例 1:

[[1,3,1],
 [1,5,1],
 [4,2,1]]
根据上面的数组，返回 7. 因为路径 1→3→1→1→1 总和最小。



裸题

package minimum_path_sum

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				matrix[i][j] = grid[0][0]
			} else if i == 0 {
				matrix[i][j] = grid[i][j] + matrix[i][j-1]
			} else if j == 0 {
				matrix[i][j] = grid[i][j] + matrix[i-1][j]
			} else {
				matrix[i][j] = matrix[i][j-1]
				if matrix[i][j] > matrix[i-1][j] {
					matrix[i][j] = matrix[i-1][j]
				}
				matrix[i][j] += grid[i][j]
			}
		}
	}
	return matrix[m-1][n-1]
}
