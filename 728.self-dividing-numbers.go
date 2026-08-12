/*
 * @lc app=leetcode id=728 lang=golang
 *
 * [728] Self Dividing Numbers
 */

// @lc code=start
func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)
	for i := left; i <= right; i++ {
		if selfDivide(i) {
			res = append(res, i)
		}
	}
	return res
}

func selfDivide(n int) bool {
	orig := n
	for n > 0 {
		d := n % 10
		if d == 0 || orig%d != 0 {
			return false
		}
		n /= 10
	}
	return true

}

// @lc code=end

