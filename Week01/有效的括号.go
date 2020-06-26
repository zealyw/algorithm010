//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串

package main

import "bytes"

func main() {
	isValid("()")
}

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	if s == "" {
		return true
	}
	var charStack []rune
	for i := 0; i < len(s); i++ {
		if bytes.ContainsRune([]byte{'(', '[', '{'}, rune(s[i])) {
			charStack = append(charStack, rune(s[i]))
		} else {
			if len(charStack) == 0 {
				return false
			} else {
				if (charStack[len(charStack)-1] == '(' && s[i] == ')') ||
					(charStack[len(charStack)-1] == '[' && s[i] == ']') ||
					(charStack[len(charStack)-1] == '{' && s[i] == '}') {
					charStack = charStack[:len(charStack)-1]
				} else {
					return false
				}
			}
		}
	}
	return len(charStack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
