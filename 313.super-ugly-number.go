/*
 * @lc app=leetcode id=313 lang=golang
 *
 * [313] Super Ugly Number
 */

// @lc code=start
func nthSuperUglyNumber(n int, primes []int) int {
	k := len(primes)
	incr := make([]int, k)
	ugly := make([]int, n)
	ugly[0] = 1

	for i := 1; i < n; i++ {
		v := int(^uint(0) >> 1)

		for j := 0; j < k; j++ {
			val := ugly[incr[j]] * primes[j]
			if val < v {
				v = val
			}
		}

		ugly[i] = v

		for j := 0; j < k; j++ {
			if ugly[incr[j]]*primes[j] == v {
				incr[j]++
			}
		}
	}

	return ugly[n-1]
}

// @lc code=end

