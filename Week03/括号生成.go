//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例：
//
// 输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
//
// Related Topics 字符串 回溯算法

package main

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	ret := []string{}
	recur(0, 0, n, "", &ret)
	return ret
}

func recur(left int, right int, n int, s string, ret *[]string) {
	if right == left && right == n {
		(*ret) = append((*ret), s)
	}
	if left < n {
		recur(left+1, right, n, s+"(", ret)
	}
	if right < left && right < n {
		recur(left, right+1, n, s+")", ret)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
