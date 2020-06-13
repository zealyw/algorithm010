<?php
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
     * @param Integer $k
     * @return ListNode
     */
    function reverseKGroup($head, $k) {
        if($head == null) return $head;
        $i = $k;
        $cur = $head;
        while ($i>0){
            if($cur != null){
                $cur = $cur->next;
                $i--;
            }else{
                break;
            }
        }
        if($i != 0){
            return $head;
        }
        //有符合数量要求的一组
        $prev = null;
        $i = $k;
        $cur = $head;
        while ($cur && $i>0){
            $next = $cur->next;
            $cur->next = $prev;
            $prev = $cur;
            $cur = $next;
            $i--;
        }
        $head->next = $this->reverseKGroup($cur,$k);
        return $prev;


        
    }
}
//leetcode submit region end(Prohibit modification and deletion)
