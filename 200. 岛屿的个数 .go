给定 '1'（陆地）和 '0'（水）的二维网格图，计算岛屿的数量。一个岛被水包围，并且通过水平或垂直连接相邻的陆地而形成。你可以假设网格的四个边均被水包围。

示例 1:

11110
11010
11000
00000
答案: 1

示例 2:

11000
11000
00100
00011
答案: 3

并查集裸题

package number_of_islands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var ret int
	visited := map[int]bool{}
	var k int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}
			k = hash(i, j, len(grid[0]))
			if _, ok := visited[k]; ok {
				continue
			}
			ret++
			queue := []int{k}
			for len(queue) != 0 {
				top := queue[0]
				queue = queue[1:]
				row, col := top/len(grid[0]), top%len(grid[0])
				if row > 0 && grid[row-1][col] == '1' {
					k = hash(row-1, col, len(grid[0]))
					if _, ok := visited[k]; !ok {
						queue = append(queue, k)
						visited[k] = true
					}
				}
				if row < len(grid)-1 && grid[row+1][col] == '1' {
					k = hash(row+1, col, len(grid[0]))
					if _, ok := visited[k]; !ok {
						queue = append(queue, k)
						visited[k] = true
					}
				}
				if col > 0 && grid[row][col-1] == '1' {
					k = hash(row, col-1, len(grid[0]))
					if _, ok := visited[k]; !ok {
						queue = append(queue, k)
						visited[k] = true
					}
				}
				if col < len(grid[0])-1 && grid[row][col+1] == '1' {
					k = hash(row, col+1, len(grid[0]))
					if _, ok := visited[k]; !ok {
						queue = append(queue, k)
						visited[k] = true
					}
				}
			}
		}
	}
	return ret
}

func hash(row, col, length int) int {
	return row*length + col
}
