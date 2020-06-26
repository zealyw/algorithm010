//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Mar
//cos 贡献此图。
//
// 示例:
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6
// Related Topics 栈 数组 双指针

package main

import "math"

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}

//leetcode submit region begin(Prohibit modification and deletion)

func trap(height []int) int {
	area := 0
	//利用栈来减少循环 减少循环的方式是维护一个单调栈
	heightStack := make([]int, 0)
	for i := 0; i < len(height); i++ {
		for len(heightStack) != 0 && height[i] > height[heightStack[len(heightStack)-1]] {
			//当前元素比栈顶元素要大即视为右边界
			top := heightStack[len(heightStack)-1]
			heightStack = heightStack[0 : len(heightStack)-1]
			if len(heightStack) != 0 {
				minHeight := int(math.Min(float64(height[i]), float64(height[heightStack[len(heightStack)-1]]))) - height[top]
				area += (i - heightStack[len(heightStack)-1] - 1) * minHeight
			}
		}
		heightStack = append(heightStack, i)
	}
	return area
	/*//暴力法 循环了三次
	maxLeftArr := make(map[int]int, 0)
	maxRightArr := make(map[int]int, 0)
	for i := 0; i < len(height); i++ {
		maxLeftArr[i] = 0
		if max, ok := maxLeftArr[i-1]; ok {
			if max > height[i-1] {
				maxLeftArr[i] = max
			} else {
				maxLeftArr[i] = height[i-1]
			}
		}
	}
	for i := len(height) - 1; i > 0; i-- {
		maxRightArr[i] = 0
		if max, ok := maxRightArr[i+1]; ok {
			if max > height[i+1] {
				maxRightArr[i] = max
			} else {
				maxRightArr[i] = height[i+1]
			}
		}
	}
	area := 0.0
	for i := 0; i < len(height); i++ {
		tempArea := math.Min(float64(maxLeftArr[i]), float64(maxRightArr[i])) - float64(height[i])
		if tempArea > 0 {
			area += tempArea
		}
	}
	return int(area)*/
}

//leetcode submit region end(Prohibit modification and deletion)
