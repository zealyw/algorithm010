package main

func main() {
	jump([]int{2, 3, 1, 1, 4})
}
func jump(nums []int) int {
	//利用贪心 从后往前 看能跳到这一格最远的是哪个格子
	times := 0
	tail := len(nums) - 1
	for tail != 0 {
		for i := 0; i < tail; i++ {
			if nums[i]+i >= tail {
				tail = i
				times++
				break
			}
		}
	}
	return times
}
