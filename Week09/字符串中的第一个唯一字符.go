//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
//
//
// 示例：
//
// s = "leetcode"
//返回 0
//
//s = "loveleetcode"
//返回 2
//
//
//
//
// 提示：你可以假定该字符串只包含小写字母。
// Related Topics 哈希表 字符串

package main

//leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	//利用hashmap
	charMap := make(map[int32]int)
	for _, char := range s {
		charMap[char]++
	}
	for i, char := range s {
		if charMap[char] > 1 {
			continue
		}
		return i
	}
	return -1

	//leetcode上看到的比较好的操作
	var arr [26]int
	for i, k := range s {
		arr[k-'a'] = i
	}
	for i, k := range s {
		if i == arr[k-'a'] {
			return i
		} else {
			arr[k-'a'] = -1
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
