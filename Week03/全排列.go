//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
// 示例:
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
// Related Topics 回溯算法

package main

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	res := [][]int{}
	backtrack([]int{}, nums, make([]bool, len(nums)), &res)
	return res
}

func backtrack(cur []int, nums []int, used []bool, res *[][]int) {
	if len(nums) == len(cur) {
		tmp := make([]int, len(nums))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i, num := range nums {
		if !used[i] {
			used[i] = true
			cur = append(cur, num)
			backtrack(cur, nums, used, res)
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
