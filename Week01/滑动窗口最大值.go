//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//
// 返回滑动窗口中的最大值。
//
//
//
// 进阶：
//
// 你能在线性时间复杂度内解决此题吗？
//
//
//
// 示例:
//
// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
// 1 <= k <= nums.length
//
// Related Topics 堆 Sliding Window

package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	/*//双端队列的方式 主要是通过维护一个双端队列 队列的首个元素是当前窗口的最大值
	//当窗口滑动的时候，新出现的值同队列尾的值对比如果当前的值比队尾的值要大就删掉这个元素因为他不可能是一个最大值了
	//当窗口滑动的时候如果队头元素已经不在窗口中了就删掉
	dequeue, ret := make([]int,0), make([]int,0)
	for i:=0;i< len(nums);i++ {
		for len(dequeue) >0 && nums[dequeue[len(dequeue)-1]] < nums[i] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		dequeue = append(dequeue,i)
		if i>=k-1{
			if dequeue[0] == i-k {
				dequeue = dequeue[1:]
			}
			ret=append(ret, nums[dequeue[0]])
		}
	}
	return ret*/
}

//leetcode submit region end(Prohibit modification and deletion)
