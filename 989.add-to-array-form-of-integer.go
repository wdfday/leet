/*
 * @lc app=leetcode id=989 lang=golang
 *
 * [989] Add to Array-Form of Integer
 */

// @lc code=start
func addToArrayForm(num []int, k int) []int {
	res := []int{}
	carry := 0
	i := len(num) - 1

	for i >= 0 || k > 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += num[i]
			i--
		}
		if k > 0 {
			sum += k % 10
			k /= 10
		}
		res = append(res, sum%10)
		carry = sum / 10
	}

	// reverse
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}

	return res
}

// @lc code=end

