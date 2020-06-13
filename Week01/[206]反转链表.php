<?php
//反转一个单链表。
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL 
//
// 进阶: 
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？ 
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
    function reverseList($head) {
        //使用迭代的方式
        $prev = null;
        $cur = $head;
        while ($cur){
            $next = $cur->next;
            $cur->next = $prev;
            $prev = $cur;
            $cur = $next;
        }
        return $prev;


        //递归方式实现反转链表
        if($head == null || $head->next == null) return $head;
        $last = $this->reverseList($head->next);
        $head->next->next = $head;
        $head->next = null;
        return $last;
















































        //递归
        if($head == null || $head->next == null) return $head;
        $last = $this->reverseList($head->next);
        $head->next->next = $head;
        $head->next = null;
        return $last;


        //迭代
        $prev = null;
        $cur = $head;
        while ($cur){
            $next = $cur->next;
            $cur->next = $prev;
            $prev = $cur;
            $cur = $next;
        }
        return $prev;









































        //迭代
        $prev = null;
        $cur = $head;
        while ($cur){
            $next = $cur->next;
            $cur->next = $prev;
            $prev = $cur;
            $cur = $next;
        }
        return $prev;
        //递归
        //找出来退出条件
        if($head == null || $head->next == null){
            return $head;
        }
        $last = $this->reverseList($head->next);
        $head->next->next = $head;
        $head->next = null;
        return $last;


    }
}
//leetcode submit region end(Prohibit modification and deletion)
