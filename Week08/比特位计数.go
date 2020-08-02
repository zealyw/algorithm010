//给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
//
// 示例 1:
//
// 输入: 2
//输出: [0,1,1]
//
// 示例 2:
//
// 输入: 5
//输出: [0,1,1,2,1,2]
//
// 进阶:
//
//
// 给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
// 要求算法的空间复杂度为O(n)。
// 你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。
//
// Related Topics 位运算 动态规划

package main

func main() {
	countBits(4)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countBits(num int) []int {
	//暴力法从0到num遍历 判断每个数的二进制数中1的个数
	/*var res []int
	for i:=0;i<=num;i++ {
		count := 0
		count = hammingWeight(uint32(i))
		res = append(res, count)
	}
	return res*/
	//O(n)的时间复杂度
	//规律 如果一个数是奇数就是 前一个偶数的1个数加1;如果一个数是偶数 它1的个数一定和除以2的位数一样
	/*res := []int{0}
	for i:=1;i<=num;i++ {
		if i & 1 == 1 {
			//奇数
			res = append(res, res[i-1] + 1)
		}else{
			res = append(res ,res[i >> 1])
		}
	}
	return res*/
	//leetcode 大佬操作
	arr := make([]int, num+1)
	for i := 1; i <= num; i++ {
		arr[i] = arr[i&(i-1)] + 1
	}
	return arr

}
func hammingWeight(num uint32) int {
	//通过n&n-1来直接判断1的位数
	count := 0
	for num != 0 {
		count++
		num &= num - 1
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
