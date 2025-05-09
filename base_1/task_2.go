package main

import "fmt"

/*
198. 打家劫舍：你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你不触动警报装置的情况下，一夜之内能够偷窃到的最高金额。这道题可以使用动态规划的思想，通过 for 循环遍历数组，
利用 if 条件判断来决定是否选择当前房屋进行抢劫，状态转移方程为 dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])。
*/

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println("最大偷窃金额为:", rob(nums)) // 输出: 12
}

// rob 返回可以偷窃的最大金额
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// 创建dp数组，dp[i]表示前i间房屋可以偷窃的最大金额
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

// max 返回两个整数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
