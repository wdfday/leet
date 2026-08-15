/*
 * @lc app=leetcode id=3622 lang=golang
 *
 * [3622] Check Divisibility by Digit Sum and Product
 */

// @lc code=start
func checkDivisibility(n int) bool {
	sum := 0
	prod := 1
	n1 := n
	for n1 > 0 {
		digit := n1%10
		n1 /= 10
		sum += digit
		prod *= digit
	}
	return n % (sum + prod) == 0
}
// @lc code=end

