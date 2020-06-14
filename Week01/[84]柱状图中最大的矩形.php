<?php
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。 
//
// 
//
// 
//
// 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。 
//
// 
//
// 
//
// 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。 
//
// 
//
// 示例: 
//
// 输入: [2,1,5,6,2,3]
//输出: 10 
// Related Topics 栈 数组


//leetcode submit region begin(Prohibit modification and deletion)
class Solution
{

    /**
     * @param Integer[] $heights
     * @return Integer
     */
    function largestRectangleArea($heights) {
        //利用有序栈
        $stack = new SplStack();
        $stack->push(['index' => -1,'height' => -1]);
        $count = count($heights);
        $maxArea = 0;
        for($i = 0; $i< $count; $i++){
            $right = $i;
            while ($heights[$i] < $stack->top()['height']){
                $height = $stack->pop()['height'];
                $left = $stack->top()['index'];
                $maxArea = max($maxArea,($right-$left-1)*$height);
            }
            $stack->push(['index' => $i,'height' => $heights[$i]]);
        }
        while (-1 != $stack->top()['index']){
            $height = $stack->pop()['height'];
            $left = $stack->top()['index'];
            $maxArea = max($maxArea,($count-$left-1)*$height);
        }
        return $maxArea;

        //暴力法2
        $count = count($heights);
        $maxArea = 0;
        for ($mid = 0; $mid < $count; $mid++) {
            //以i为中心找到左右边界
            $left = $mid;
            $right = $mid;
            while ($right+1 < $count && $heights[$right+1] >=$heights[$mid]){
                $right++;
            }
            while ($left-1 >= 0 && $heights[$left-1] >=$heights[$mid]){
                $left--;
            }
            $maxArea = max($maxArea,($right-$left+1)*$heights[$mid]);
        }
        return $maxArea;



        //暴力法 会超时
        $count = count($heights);
        $maxArea = 0;
        for ($i = 0; $i < $count; $i++) {
            $minHeight = $heights[$i];
            $maxArea = max($maxArea, $minHeight);
            for ($j = $i + 1; $j < $count; $j++) {
                $minHeight = $minHeight > $heights[$j] ? $heights[$j] : $minHeight;
                $maxArea = max($maxArea, ($j - $i + 1) * $minHeight);
            }
        }
        return $maxArea;

    }
}
//leetcode submit region end(Prohibit modification and deletion)
