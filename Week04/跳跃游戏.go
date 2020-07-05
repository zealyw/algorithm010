//给定一个非负整数数组，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个位置。
//
// 示例 1:
//
// 输入: [2,3,1,1,4]
//输出: true
//解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
//
//
// 示例 2:
//
// 输入: [3,2,1,0,4]
//输出: false
//解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
//
// Related Topics 贪心算法 数组

package main

func main() {
	canJump([]int{2, 3, 1, 1, 4})
}

//leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	//递归 记录访问过的点 可以过用例但是会超时
	//return jump(nums,0,&map[int]bool{})
	//利用贪心，从后往前贪心，找能到最后这个点的位置
	if len(nums) == 0 {
		return false
	}
	endPoint := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+i >= endPoint {
			endPoint = i
		}
	}
	return endPoint == 0
}
func jump(nums []int, i int, visited *map[int]bool) bool {
	//终止条件
	if i >= len(nums)-1 {
		return true
	}
	if (*visited)[i] == true {
		return false
	}
	(*visited)[i] = true
	//当前层处理
	for j := 1; j <= nums[i]; j++ {
		//下一层
		if jump(nums, i+j, visited) == true {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
