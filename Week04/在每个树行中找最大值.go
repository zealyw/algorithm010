//您需要在二叉树的每一行中找到最大的值。
//
// 示例：
//
//
//输入:
//
//          1
//         / \
//        3   2
//       / \   \
//      5   3   9
//
//输出: [1, 3, 9]
//
// Related Topics 树 深度优先搜索 广度优先搜索

package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	//层序遍历然后获取这一行的最大值 遍历过程中就可以记录
	if root == nil {
		return nil
	}
	var res []int
	var curLevelQueue = []TreeNode{*root}
	for len(curLevelQueue) != 0 {
		curQueueLen := len(curLevelQueue)
		max := math.MinInt64
		for i := 0; i < curQueueLen; i++ {
			if max < curLevelQueue[i].Val {
				max = curLevelQueue[i].Val
			}
			if curLevelQueue[i].Left != nil {
				curLevelQueue = append(curLevelQueue, *curLevelQueue[i].Left)
			}
			if curLevelQueue[i].Right != nil {
				curLevelQueue = append(curLevelQueue, *curLevelQueue[i].Right)
			}
		}
		res = append(res, max)
		curLevelQueue = curLevelQueue[curQueueLen:]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
