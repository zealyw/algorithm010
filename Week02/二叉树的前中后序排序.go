//给定一个二叉树，返回它的 前序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,2,3]
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	//根左右
	if root == nil {
		return []int{}
	}
	path := append([]int{root.Val}, preorderTraversal(root.Left)...)
	path = append(path, preorderTraversal(root.Right)...)
	return path
}

func inorderTraversal(root *TreeNode) []int {
	//左根右
	if root == nil {
		return []int{}
	}
	path := append(inorderTraversal(root.Left), root.Val)
	path = append(path, inorderTraversal(root.Right)...)
	return path
}
func postorderTraversal(root *TreeNode) []int {
	//左右根
	if root == nil {
		return []int{}
	}
	path := append(postorderTraversal(root.Left), postorderTraversal(root.Right)...)
	path = append(path, root.Val)
	return path
}

//leetcode submit region end(Prohibit modification and deletion)
