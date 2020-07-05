//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
//
//
//
// 每次转换只能改变一个字母。
// 转换过程中的中间单词必须是字典中的单词。
//
//
// 说明:
//
//
// 如果不存在这样的转换序列，返回 0。
// 所有单词具有相同的长度。
// 所有单词只由小写字母组成。
// 字典中不存在重复的单词。
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
//
//
// 示例 1:
//
// 输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出: 5
//
//解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//     返回它的长度 5。
//
//
// 示例 2:
//
// 输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: 0
//
//解释: endWord "cog" 不在字典中，所以无法进行转换。
// Related Topics 广度优先搜索

package main

func main() {
	ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
}

//leetcode submit region begin(Prohibit modification and deletion)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if idxof(endWord, wordList) == -1 {
		return 0
	}
	step := 0
	used := make([]bool, len(wordList))
	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	for len(startQueue) > 0 {
		step++
		l := len(startQueue)
		for i := 0; i < l; i++ {
			for k := 0; k < len(endQueue); k++ {
				if hasOneDiff(startQueue[i], endQueue[k]) {
					return step + 1
				}
			}
			for j, w := range wordList {
				if !used[j] && hasOneDiff(startQueue[i], w) {
					startQueue = append(startQueue, w)
					used[j] = true
				}
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}
	return 0
}

func hasOneDiff(s string, w string) bool {
	diffCount := 0
	for i := 0; i < len(w); i++ {
		if s[i] != w[i] {
			diffCount++
			if diffCount > 1 {
				return false
			}
		}
	}
	return diffCount == 1
}

func idxof(targetWord string, wordList []string) int {
	for index, word := range wordList {
		if word == targetWord {
			return index
		}
	}
	return -1
}

/*func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) != len(endWord) || len(wordList) == 0 {
		return 0
	}
	times := 1
	wordMap := map[string]bool{}
	for _, word := range wordList {
		wordMap[word] = true
	}
	if wordMap[endWord] != true {
		return 0
	}
	used := map[string]bool{}
	curLevelQueue := []string{beginWord}
	for len(curLevelQueue) != 0 {
		curLevelQueueLen := len(curLevelQueue)
		for i := 0; i < curLevelQueueLen; i++ {
			word := curLevelQueue[i]
			if word == endWord {
				return times
			}
			for index := 0; index < len(word); index++ {
				for j := 0; j < 26; j++ {
					newWord := word[:index] + string(byte('a'+j)) + word[index+1:]
					if _, ok := used[newWord]; !ok && wordMap[newWord] == true{
						curLevelQueue = append(curLevelQueue, newWord)
					}
					used[newWord] = true
				}
			}
		}
		curLevelQueue = curLevelQueue[curLevelQueueLen:]
		times++
	}
	return 0
}*/

//leetcode submit region end(Prohibit modification and deletion)
