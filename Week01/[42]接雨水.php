<?php
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 
//
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Mar
//cos 贡献此图。 
//
// 示例: 
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6 
// Related Topics 栈 数组 双指针


//leetcode submit region begin(Prohibit modification and deletion)
class Solution
{

    /**
     * @param Integer[] $height
     * @return Integer
     */
    function trap($height) {
        $count = count($height);
        $ans = 0;
        //使用双指针
        $right = $count-1;
        $left = 0;
        $leftMax = 0;
        $rightMax = 0;
        while ($right>$left){
            if($height[$right] > $height[$left]){
                if($height[$left] >= $leftMax){
                    $leftMax = $height[$left];
                }else{
                    $ans += $leftMax-$height[$left];
                }
                $left++;
            }elseif ($height[$right] <= $height[$left]){
                if($height[$right] >= $rightMax){
                    $rightMax = $height[$right];
                }else{
                    $ans += $rightMax-$height[$right];
                }
                $right--;
            }
        }
        return $ans;
        //再优化 利用栈
        $stack = new SplStack();
        for($i = 0;$i<$count;$i++){
            while(!$stack->isEmpty() && $height[$i]>$height[$stack->top()]){
                $pop = $stack->pop();
                while (!$stack->isEmpty() && $height[$pop] == $height[$stack->top()]){
                    $stack->pop();
                }
                if(!$stack->isEmpty()){
                    //这就是左边界了
                    $top = $stack->top();
                    $ans += ($i-$top -1) * (min($height[$i],$height[$top]) - $height[$pop]);
                }
            }
            $stack->push($i);
        }
        return $ans;
        //优化： 先将左右两边的最大值都记录下来
        $maxRightArr = [];
        $maxLeftArr = [];
        for ($i = 0; $i < $count; $i++) {
            $maxLeftArr[$i] = isset($maxLeftArr[$i - 1]) && $maxLeftArr[$i - 1] > $height[$i] ? $maxLeftArr[$i - 1] : $height[$i];
        }
        for ($i = $count - 1; $i >= 0; $i--) {
            $maxRightArr[$i] = isset($maxRightArr[$i + 1]) && $maxRightArr[$i + 1] > $height[$i] ? $maxRightArr[$i + 1] : $height[$i];
        }
        for ($i = 1; $i < $count - 1; $i++) {
            $ans += min($maxLeftArr[$i], $maxRightArr[$i]) - $height[$i];
        }
        return $ans;

        //暴力法 存在的问题是为了找最大值每次在左右两遍都扫了一次
        for ($i = 1; $i < $count - 1; $i++) {
            $maxRight = $maxLeft = 0;
            for ($j = $i; $j < $count; $j++) {
                $maxRight = max($maxRight, $height[$j]);
            }
            for ($j = $i; $j >= 0; $j--) {
                $maxLeft = max($maxLeft, $height[$j]);
            }
            $ans += min($maxLeft, $maxRight) - $height[$i];
        }
        return $ans;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
