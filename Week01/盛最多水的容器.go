//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i,
//ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器，且 n 的值至少为 2。
//
//
//
//
//
// 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
//
//
// 示例：
//
// 输入：[1,8,6,2,5,4,8,3,7]
//输出：49
// Related Topics 数组 双指针

package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	maxArea := 0.0
	right := len(height) - 1
	left := 0
	//简化代码 简化的点在于go自带的min函数只能比float
	for left < right {
		maxArea = math.Max(maxArea, math.Min(float64(height[right]), float64(height[left]))*float64(right-left))
		if height[right] > height[left] {
			left += 1
		} else {
			right -= 1
		}
	}
	/*for left<right {
		minHeight := 0
		if height[right] > height[left]{
			minHeight = height[left]
			left+=1
		}else{
			minHeight = height[right]
			right -=1
		}
		if (right-left+1)* minHeight > maxArea{
			maxArea = (right-left+1)* minHeight
		}
	}*/
	return int(maxArea)
}

//leetcode submit region end(Prohibit modification and deletion)
