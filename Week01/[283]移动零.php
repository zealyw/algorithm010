<?php
//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
//
// Related Topics 数组 双指针


//leetcode submit region begin(Prohibit modification and deletion)
class Solution
{

    /**
     * @param Integer[] $nums
     * @return NULL
     */
    function moveZeroes(&$nums) {

        //循环一次数组，所有循环的元素非0的元素在前，0元素跟在后面
        //设置一个变量来记已循环0元素里第一个的index，发现遍历到的元素不是0的时候就与第一个0元素交换位置，然后第一个零元素的位置后移
        //方案1:复杂度为O(n)循环遍历一次数组
        $firstZeroIndex = -1;
        for ($i = 0; $i < count($nums); $i++) {
            $firstZeroIndex == -1 && $nums[$i] == 0 && $firstZeroIndex = $i;
            if ($firstZeroIndex != -1 && $nums[$i] != 0) {
                $nums[$firstZeroIndex++] = $nums[$i];
                $nums[$i] = 0;
            }
        }


        /**************************************************************
         * //使用PHP的函数来做的话1.获取元素有多少个0 2.过滤掉数组的0 3.补充原有个数的0
         * $numsCount = count($nums);
         * $nums = array_filter($nums);
         * $filterNumsCount = count($nums);
         * $numsCount-$filterNumsCount >0 && $nums = array_merge($nums,array_fill(0,$numsCount-$filterNumsCount,0));
         **************************************************************/


        return true;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
