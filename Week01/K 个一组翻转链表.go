//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//
//
// 示例：
//
// 给你这个链表：1->2->3->4->5
//
// 当 k = 2 时，应当返回: 2->1->4->3->5
//
// 当 k = 3 时，应当返回: 3->2->1->4->5
//
//
//
// 说明：
//
//
// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
// Related Topics 链表

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	i, cur := k, head
	for ; i > 0; cur, i = cur.Next, i-1 {
		if cur == nil {
			return head
		}
	}
	//存在k个元素可以翻转
	cur = head
	prev := &ListNode{Val: 0, Next: head}
	i = k
	for cur != nil && i > 0 {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		i--
	}
	head.Next = reverseKGroup(cur, k)
	return prev
}

//leetcode submit region end(Prohibit modification and deletion)
