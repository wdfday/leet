/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start
func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			sq := j * j
			dp[i] = min(dp[i], dp[i-sq]+1)
		}
	}
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func numSquares(n int) int {
// 	visited := make([]bool, n+1)
// 	queue := []int{n}
// 	step := 0

// 	for len(queue) > 0 {
// 		size := len(queue)
// 		step++

// 		for i := 0; i < size; i++ {
// 			cur := queue[0]
// 			queue = queue[1:]

// 			for j := 1; j*j <= cur; j++ {
// 				next := cur - j*j
// 				if next == 0 {
// 					return step
// 				}
// 				if !visited[next] {
// 					visited[next] = true
// 					queue = append(queue, next)
// 				}
// 			}
// 		}
// 	}

// 	return n
// }

// @lc code=end

