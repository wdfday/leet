/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	result := 1.0
	currentProduct := x
	for i := n; i > 0; i /= 2 {
		if i%2 == 1 {
			result *= currentProduct
		}
		currentProduct *= currentProduct
	}
	return result
}

// @lc code=end

