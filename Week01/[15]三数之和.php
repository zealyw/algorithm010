<?php
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例： 
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
// 
// Related Topics 数组 双指针


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function threeSum($nums) {
        //1.双指针方法
        $count = count($nums);
        sort($nums);
        $result = [];
        for($i=0;$i<$count-2;$i++){
            $target = -$nums[$i];
            if($i>0 && $nums[$i] == $nums[$i-1]){
                continue;
            }
            for($j=$i+1,$k=$count-1;$j<$k;){
                if($target > ($nums[$j]+$nums[$k])){
                    $j++;
                }elseif ($target < ($nums[$j] + $nums[$k])){
                    $k--;
                }else{
                    $result[] = [$nums[$i],$nums[$j],$nums[$k]];
                    //去重
                    while($j<$k && $nums[$j] == $nums[++$j]);
                    while($j<$k && $nums[$k] == $nums[--$k]);
                }
            }
        }

        return $result;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
