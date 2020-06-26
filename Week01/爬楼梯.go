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

package main

//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	//到第n级台阶有两种方式，一种是从第n-1级台阶跨一步上来，一种是从第n-2级台阶跨两级上来
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	fib1, fib2, fib3 := 1, 2, 3
	for i := 3; i <= n; i++ {
		fib3 = fib1 + fib2
		fib1 = fib2
		fib2 = fib3
	}
	return fib3
}

//leetcode submit region end(Prohibit modification and deletion)
