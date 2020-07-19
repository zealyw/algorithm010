package main

import "sort"

func main() {
	fourSum([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0)
}
func fourSum(nums []int, target int) [][]int {
	var ret [][]int
	sort.Ints(nums)
	for firstNum := 0; firstNum < len(nums)-3; firstNum++ {
		if firstNum > 0 && nums[firstNum] == nums[firstNum-1] {
			continue
		}
		for secondNum := firstNum + 1; secondNum < len(nums)-2; {
			if secondNum > firstNum+1 && nums[secondNum] == nums[secondNum-1] {
				continue
			}
			newTarget := target - nums[firstNum] - nums[secondNum]
			thirdNum, fourthNum := secondNum+1, len(nums)-1
			for thirdNum < fourthNum {
				if (nums[thirdNum] + nums[fourthNum]) > newTarget {
					fourthNum--
				} else if (nums[thirdNum] + nums[fourthNum]) < newTarget {
					thirdNum++
				} else {
					ret = append(ret, []int{nums[firstNum], nums[secondNum], nums[thirdNum], nums[fourthNum]})
					for thirdNum += 1; thirdNum < len(nums) && nums[thirdNum] == nums[thirdNum-1]; thirdNum++ {
						continue
					}
					for fourthNum -= 1; fourthNum >= 0 && nums[fourthNum] == nums[fourthNum+1]; fourthNum-- {
						continue
					}
				}
			}
			for secondNum += 1; secondNum < len(nums) && nums[secondNum] == nums[secondNum-1]; secondNum++ {
				continue
			}
		}
	}
	return ret
}
