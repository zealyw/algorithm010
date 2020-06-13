<?php
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


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */
class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    function swapPairs($head) {

        //递归法
        if($head == null || $head->next  == null){
            return $head;
        }
        $next = $head->next;
        $head->next = $this->swapPairs($head->next->next);
        $next->next = $head;
        return $next;

        //迭代法
        if($head == null || $head->next == null){
            return $head;
        }
        $listNode = new ListNode(null);
        $listNode->next = $head;
        $cur = $listNode;
        while ($cur->next != null && $cur->next->next != null){
            $first = $cur->next;
            $second = $cur->next->next;
            $cur->next = $second;
            $first->next = $second->next;
            $second->next = $first;
            $cur = $cur->next->next;
        }
        return $listNode->next;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
