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
//[
//['1','1','1','1','0'],
//['1','1','0','1','0'],
//['1','1','0','0','0'],
//['0','0','0','0','0']
//]
//输出: 1
//
//
// 示例 2:
//
// 输入:
//[
//['1','1','0','0','0'],
//['1','1','0','0','0'],
//['0','0','1','0','0'],
//['0','0','0','1','1']
//]
//输出: 3
//解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
//
// Related Topics 深度优先搜索 广度优先搜索 并查集

package main

func main() {
	numIslands([][]byte{[]byte{'1', '1', '1', '1', '0'}, []byte{'1', '1', '0', '1', '0'}, []byte{'1', '1', '0', '0', '0'}, []byte{'0', '0', '0', '0', '0'}})
}

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	//dfs方式
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				result++
				setVisited(i, j, grid)
			}
		}
	}
	return result
}
func setVisited(i int, j int, grid [][]byte) {
	grid[i][j] = '0'
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		setVisited(i+1, j, grid)
	}
	if i-1 >= 0 && grid[i-1][j] == '1' {
		setVisited(i-1, j, grid)
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
		setVisited(i, j+1, grid)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		setVisited(i, j-1, grid)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
