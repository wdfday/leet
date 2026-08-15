/*
 * @lc app=leetcode id=3697 lang=golang
 *
 * [3697] Compute Decimal Representation
 */

// @lc code=start
func decimalRepresentation(n int) []int {
    res := []int{}
	k := 1
	for n > 0 {
		v := n%10
		n /= 10
		if v > 0 {
			res = append(res, v * k)
		}
		k *= 10
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
// @lc code=end

