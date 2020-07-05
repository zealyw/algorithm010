//给定一个可包含重复数字的序列，返回所有不重复的全排列。
//
// 示例:
//
// 输入: [1,1,2]
//输出:
//[
//  [1,1,2],
//  [1,2,1],
//  [2,1,1]
//]
// Related Topics 回溯算法

package main

func main() {
	permuteUnique([]int{1, 1, 2})
}

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	res, cur := [][]int{}, []int{}
	backtrack(cur, nums, make([]bool, len(nums)), &res)
	return res
}
func backtrack(cur []int, nums []int, visited []bool, res *[][]int) {
	if len(cur) == len(nums) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	used := map[int]bool{}
	for i, num := range nums {
		if value, ok := used[num]; (ok && value) || visited[i] {
			continue
		}
		used[num] = true
		visited[i] = true
		cur = append(cur, num)
		backtrack(cur, nums, visited, res)
		cur = cur[:len(cur)-1]
		visited[i] = false
	}
}

//leetcode submit region end(Prohibit modification and deletion)
