<?php
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
// 示例 1：
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//
// 示例 2：
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
//
// Related Topics 动态规划


//leetcode submit region begin(Prohibit modification and deletion)
class Solution
{

    /**
     * @param Integer $n
     * @return Integer
     */
    function climbStairs($n) {
        //思路找重复操作 第n个台阶上去的方法要么从n-1层 上去1层，要么在第n-2层的时候一次两个台阶 转换之后就是 f(n) = f(n-1)+f(n-2);
        /*
        //直接套路公式，会超时 时间复杂度是2^n
        if($n==1 || $n==0){
            return 1;
        }
        return $this->climbStairs($n-1) + $this->climbStairs($n-2);
        */
        /*
        //改进方案 减少重复计算 时间复杂度为O(n)
        $fib = [
                0 => 1,
                1 => 1,
        ];
        for($i=1;$i<=$n;++$i){
            $fib[$i] = $fib[$i-1] + $fib[$i-2];
        }
        return $fib[$n];
        */
        //关注题目自身 只需要求fib(n)所以只需要一直留存f(n-1)和f(n-2)
        if ($n <= 2)
            return $n;
        $fib1 = 1;
        $fib2 = 2;
        $fib3 = 3;
        for ($i = 3; $i <= $n; ++$i) {
            $fib3 = $fib1 + $fib2;
            $fib1 = $fib2;
            $fib2 = $fib3;
        }
        return $fib3;


    }
}
//leetcode submit region end(Prohibit modification and deletion)
