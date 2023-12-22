package leetMinCostClimbingStairs

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+2)
	dp[0] = 0
	var i int
	for i = 1; i <= 2 && i <= n; i++ {
		dp[i] = cost[i-1]
	}
	for i = 3; i <= n; i++ {
		dp[i] = cost[i-1] + min(dp[i-1], dp[i-2])
	}
	dp[n+1] = min(dp[n], dp[n-1])
	return dp[n+1]
}
