/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 */

// @lc code=start
func getPermutation(n int, k int) string {

	factorials := make([]int, n)
	factorials[0] = 1
	for i := 1; i < n; i++ {
		factorials[i] = factorials[i-1] * i
	}

	k--
	result := make([]byte, n)
	available := make([]bool, n)

	for i := 0; i < n; i++ {
		factorial := factorials[n-1-i]
		index := k / factorial
		k %= factorial

		count := 0
		for j := 0; j < n; j++ {
			if !available[j] {
				if count == index {
					result[i] = byte(j + 1 + '0')
					available[j] = true
					break
				}
				count++
			}
		}
	}

	return string(result)

}

// @lc code=end

