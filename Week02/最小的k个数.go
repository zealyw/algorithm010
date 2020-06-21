//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
//
//
// 示例 1：
//
// 输入：arr = [3,2,1], k = 2
//输出：[1,2] 或者 [2,1]
//
//
// 示例 2：
//
// 输入：arr = [0,1,2,1], k = 1
//输出：[0]
//
//
//
// 限制：
//
//
// 0 <= k <= arr.length <= 10000
// 0 <= arr[i] <= 10000
//
// Related Topics 堆 分治算法

package main

import "container/heap"

func main() {
	getLeastNumbers([]int{3, 2, 1}, 2)
}

//leetcode submit region begin(Prohibit modification and deletion)

type numsHeap []int

func (n numsHeap) Len() int {
	return len(n)
}

func (n numsHeap) Less(i, j int) bool {
	return n[i] > n[j]
}

func (n numsHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *numsHeap) Push(x interface{}) {
	*n = append(*n, x.(int))
}

func (n *numsHeap) Pop() interface{} {
	val := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return val
}
func getLeastNumbers(arr []int, k int) []int {
	//维护一个k个数的最大堆，每次进来的数都和堆顶的数比较，如果比这个数要大就替换掉这个数
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	initHeap := make(numsHeap, 0, k)
	heap.Init(&initHeap)
	for _, v := range arr {
		if len(initHeap) < k {
			heap.Push(&initHeap, v)
		} else if len(initHeap) == k {
			if v < initHeap[0] {
				heap.Pop(&initHeap)
				heap.Push(&initHeap, v)
			}
		}
	}
	return initHeap
}

//leetcode submit region end(Prohibit modification and deletion)
