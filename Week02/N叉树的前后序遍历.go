//给定一个 N 叉树，返回其节点值的前序遍历。
//
// 例如，给定一个 3叉树 :
//
//
//
//
//
//
//
// 返回其前序遍历: [1,3,5,6,2,4]。
//
//
//
// 说明: 递归法很简单，你可以使用迭代法完成此题吗? Related Topics 树

package main

type Node struct {
	Val      int
	Children []*Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	path := []int{root.Val}
	for _, child := range root.Children {
		path = append(path, preorder(child)...)
	}
	return path
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	path := []int{}
	for _, child := range root.Children {
		path = append(path, postorder(child)...)
	}
	path = append(path, root.Val)
	return path
}

//leetcode submit region end(Prohibit modification and deletion)
