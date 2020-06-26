//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
//
//
// 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
//
//
//
//
//
// 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
//
//
//
// 示例:
//
// 输入: [2,1,5,6,2,3]
//输出: 10
// Related Topics 栈 数组

package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
	area := 0.0
	//维护一个单调递增的栈
	heightStack := make([]int, 0)
	//为了方便在头尾加上为0的柱子
	heights = append(append([]int{0}, heights...), 0)
	for index, height := range heights {
		for len(heightStack) != 0 && height < heights[heightStack[len(heightStack)-1]] {
			top := heights[heightStack[len(heightStack)-1]]
			heightStack = heightStack[0 : len(heightStack)-1]
			area = math.Max(area, float64((index-heightStack[len(heightStack)-1]-1)*top))
		}
		heightStack = append(heightStack, index)
	}
	return int(area)
	/*
		//以当前节点为最最低点 往左往右找比他高的
		area := 0.0
		for i:= 0; i< len(heights);i++{
			l:=i-1
			for ;l>=0; l--{
				if heights[i]>heights[l]{
					break
				}
			}
			r:=i+1
			for ;r<len(heights); r++{
				if heights[i]>heights[r]{
					break
				}
			}
			area = math.Max(area,float64(heights[i] * (r-l-1)))
		}
		return int(area)
	*/
}

//leetcode submit region end(Prohibit modification and deletion)
