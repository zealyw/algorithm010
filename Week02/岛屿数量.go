//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1:
//
// 输入:
//11110
//11010
//11000
//00000
//输出: 1
//
//
// 示例 2:
//
// 输入:
//11000
//11000
//00100
//00011
//输出: 3
//解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
//
// Related Topics 深度优先搜索 广度优先搜索 并查集

package main

func main() {
	println(numIslands([][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	nr := len(grid)
	if nr == 0 {
		return 0
	}
	nc := len(grid[0])
	result := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				result += 1
				//然后需要所有与这个陆地相连的都置访问过
				dfs(&grid, r, c, nr, nc)
			}
		}
	}
	return result
}

func dfs(grid *[][]byte, r int, c int, nr int, nc int) {
	(*grid)[r][c] = '0'
	if r-1 >= 0 && (*grid)[r-1][c] == '1' {
		dfs(grid, r-1, c, nr, nc)
	}
	if r+1 < nr && (*grid)[r+1][c] == '1' {
		dfs(grid, r+1, c, nr, nc)
	}
	if c-1 >= 0 && (*grid)[r][c-1] == '1' {
		dfs(grid, r, c-1, nr, nc)
	}
	if c+1 < nc && (*grid)[r][c+1] == '1' {
		dfs(grid, r, c+1, nr, nc)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
