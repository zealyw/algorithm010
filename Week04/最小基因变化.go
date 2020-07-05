//一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
//
// 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
//
// 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
//
// 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
//
// 现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变
//化次数。如果无法实现目标变化，请返回 -1。
//
// 注意:
//
//
// 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
// 所有的目标基因序列必须是合法的。
// 假定起始基因序列与目标基因序列是不一样的。
//
//
// 示例 1:
//
//
//start: "AACCGGTT"
//end:   "AACCGGTA"
//bank: ["AACCGGTA"]
//
//返回值: 1
//
//
// 示例 2:
//
//
//start: "AACCGGTT"
//end:   "AAACGGTA"
//bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
//
//返回值: 2
//
//
// 示例 3:
//
//
//start: "AAAAACCC"
//end:   "AACCCCCC"
//bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
//
//返回值: 3
//
//

package main

func main() {
	println(minMutation("AACCTTGG", "AATTCCGG", []string{"AATTCCGG", "AACCTGGG", "AACCCCGG", "AACCTACC"}))
}

//leetcode submit region begin(Prohibit modification and deletion)

func minMutation(start string, end string, bank []string) int {
	var bankMap = map[string]bool{}
	var changes = []string{"A", "C", "G", "T"}

	for _, value := range bank {
		bankMap[value] = true
	}
	if len(start) != len(end) || !isInBank(end, bankMap) {
		return -1
	}
	//采用广度优先的方式
	return bfs(0, start, end, bankMap, changes)
}
func bfs(changeTimes int, start string, end string, bankMap map[string]bool, changes []string) int {
	if start == end {
		return changeTimes
	}
	var curLevelQueue []string
	var nextLevelQueue []string
	var visited = make(map[string]bool)
	curLevelQueue = append(curLevelQueue, start)
	for len(curLevelQueue) != 0 {
		node := curLevelQueue[0]
		visited[node] = true
		for i := 0; i < len(node); i++ {
			for _, change := range changes {
				newNode := node[:i] + change + node[i+1:]
				if _, ok := visited[newNode]; !ok && isInBank(newNode, bankMap) {
					if newNode == end {
						return changeTimes + 1
					} else {
						visited[newNode] = true
						nextLevelQueue = append(nextLevelQueue, newNode)
					}
				}
			}
		}
		curLevelQueue = curLevelQueue[1:]
		if len(curLevelQueue) == 0 {
			curLevelQueue = nextLevelQueue
			nextLevelQueue = []string{}
			changeTimes++
		}
	}
	return -1
}
func isInBank(str string, bankMap map[string]bool) bool {
	if _, ok := bankMap[str]; ok {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
