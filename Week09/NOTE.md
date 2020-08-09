## 动态规划
* 分解出最近循环子问题
* 分治 + 最优子结构
* 顺推形式：动态递推
### 递推模板
```
function DP():
	dp = [][] #(二维情况) 状态数组
	
	for i=1 .. M:
		for j=1 .. N:
			dp[i][j] = _Function( dp[i'][j']... ) #状态转移方程，难（简单加减，最值，k个状态最值）
	return dp[M][N]
}
```
股票买卖问题参考：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/
## 字符串匹配算法
* 暴力法
* Rabin-Karp算法
* KMP算法
* Sunday算法
