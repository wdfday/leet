/*
 * @lc app=leetcode id=3726 lang=golang
 *
 * [3726] Remove Zeros in Decimal Representation
 */

// @lc code=start
func removeZeros(n int64) int64 {
	s := strconv.Itoa(int(n))
	res := int64(0)
	for i := 0; i < len(s); i++{
		if s[i] != '0' {
			res = res * 10 + int64(s[i] - '0')
		}
	}
    return res
}
// @lc code=end

