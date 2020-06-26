//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例:
//
// 给定 1->2->3->4, 你应该返回 2->1->4->3.
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
func swapPairs(head *ListNode) *ListNode {
	//迭代实现
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	ret := head.Next
	prev := &ListNode{Val: 0, Next: head}
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur
		prev.Next = next
		prev = cur
		cur = cur.Next
	}
	return ret
	/*
		//递归实现
		if head == nil || head.Next == nil {
			return head
		}
		first,second := head,head.Next
		first.Next = swapPairs(second.Next)
		second.Next = first
		return second
	*/
}

//leetcode submit region end(Prohibit modification and deletion)
