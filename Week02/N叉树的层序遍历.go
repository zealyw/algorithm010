//给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
//
// 例如，给定一个 3叉树 :
//
//
//
//
//
//
//
// 返回其层序遍历:
//
// [
//     [1],
//     [3,2,4],
//     [5,6]
//]
//
//
//
//
// 说明:
//
//
// 树的深度不会超过 1000。
// 树的节点总数不会超过 5000。
// Related Topics 树 广度优先搜索

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*Node{}
	level := 0
	result = append(result, []int{})
	queue = append(queue, root)
	for queue != nil {
		size := len(queue)
		for i := 0; i < size; i++ {
			result[level] = append(result[level], queue[i].Val)
			for _, child := range queue[i].Children {
				queue = append(queue, child)
			}
		}
		queue = queue[size:]
		level++
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
