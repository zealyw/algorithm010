//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例：
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
// Related Topics 数组 双指针

package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	var ret [][]int
	//先对数组进行排序 然后利用双指针来计算两数合，需要注意的是去重的判断
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if (nums[left] + nums[right]) > -nums[i] {
				right--
			} else if (nums[left] + nums[right]) < -nums[i] {
				left++
			} else {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left += 1; left < len(nums) && nums[left] == nums[left-1]; left++ {
					continue
				}
				for right -= 1; right >= 0 && nums[right] == nums[right+1]; right-- {
					continue
				}
			}
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
