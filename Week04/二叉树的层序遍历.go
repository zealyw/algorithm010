//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例：
//二叉树：[3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层次遍历结果：
//
// [
//  [3],
//  [9,20],
//  [15,7]
//]
//
// Related Topics 树 广度优先搜索

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	//层序遍历的话使用bfs 一层一层来扫 关键点是要记住当前节点是哪一层的
	var res [][]int
	var curLevelQueue, nextLevelQueue []TreeNode
	var levelNodeValue []int
	if root == nil {
		return res
	}
	curLevelQueue = append(curLevelQueue, *root)
	for len(curLevelQueue) != 0 {
		node := curLevelQueue[0]
		levelNodeValue = append(levelNodeValue, node.Val)
		if node.Left != nil {
			nextLevelQueue = append(nextLevelQueue, *node.Left)
		}
		if node.Right != nil {
			nextLevelQueue = append(nextLevelQueue, *node.Right)
		}
		curLevelQueue = curLevelQueue[1:]
		if len(curLevelQueue) == 0 {
			curLevelQueue = nextLevelQueue
			nextLevelQueue = []TreeNode{}
			if len(levelNodeValue) != 0 {
				res = append(res, levelNodeValue)
				levelNodeValue = []int{}
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
