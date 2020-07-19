//一条包含字母 A-Z 的消息通过以下方式进行了编码：
//
// 'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//
//
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。
//
// 示例 1:
//
// 输入: "12"
//输出: 2
//解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
//
//
// 示例 2:
//
// 输入: "226"
//输出: 3
//解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
//
// Related Topics 字符串 动态规划

package main

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)
func numDecodings(s string) int {
	//使用dfs回溯 不过好慢
	/*sum := 0
	dfsHelper(&sum,s)
	return sum*/
	//使用dp 有点类似硬币问题 一次走一步 一次走两步
	length := len(s)
	if length == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, length+1)
	dp[0] = 1
	dp[1] = 1
	var i int
	var num int
	for i = 2; i <= length; i++ {
		if s[i-1] != '0' {
			num, _ = strconv.Atoi(s[i-2 : i])
			if num <= 26 && num >= 10 {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			num, _ = strconv.Atoi(s[i-2 : i])
			if num > 26 || num == 0 {
				return 0
			} else {
				dp[i] = dp[i-2]
			}
		}
	}
	return dp[length]
}

func dfsHelper(sum *int, s string) {
	if len(s) == 0 {
		*sum++
		return
	}
	if (s[0] - '0') == 0 {
		return
	}
	dfsHelper(sum, s[1:])
	if len(s) > 1 {
		m := (s[0]-'0')*10 + (s[1] - '0')
		if m >= 1 && m <= 26 {
			dfsHelper(sum, s[2:])
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
