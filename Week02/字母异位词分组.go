//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串

package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
type bytes []byte

// 自定义排序规则
func (b bytes) Len() int {
	return len(b)
}
func (b bytes) Less(i, j int) bool {
	return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {
	//思路：遍历所有字符串 看排除后的字符串是否出现过
	ret := [][]string{}
	stringMap := make(map[string]int)
	for _, str := range strs {
		strBytes := bytes(str)
		sort.Sort(strBytes)
		sortStr := string(strBytes)
		if idx, ok := stringMap[sortStr]; ok {
			ret[idx] = append(ret[idx], str)
		} else {
			stringMap[sortStr] = len(ret)
			ret = append(ret, []string{str})
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
