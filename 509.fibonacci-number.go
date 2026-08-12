/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(n int) int {
	f := make([]int, n+1)
	for i := range n + 1 {
		switch i {
		case 0:
			f[i] = 0
		case 1:
			f[i] = 1
		default:
			f[i] = f[i-1] + f[i-2]
		}
	}
	return f[n]
}

// @lc code=end

