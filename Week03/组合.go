//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
// 示例:
//
// 输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
// Related Topics 回溯算法

package main

//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	if n == 0 || k == 0 {
		return [][]int{}
	}
	var cur []int
	var res [][]int
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	getCombine(nums, 0, k, cur, &res)
	return res
}

func getCombine(nums []int, i int, k int, cur []int, res *[][]int) {
	if len(cur) == k {
		tmp := make([]int, k)
		copy(tmp, cur)
		*res = append(*res, tmp)
	}
	for ; i < len(nums); i++ {
		cur = append(cur, nums[i])
		getCombine(nums, i+1, k, cur, res)
		cur = cur[:len(cur)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
