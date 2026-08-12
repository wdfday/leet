/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */

// @lc code=start
func readBinaryWatch(turnedOn int) []string {
	res := []string{}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if countBits(i)+countBits(j) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return res

}

func countBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// @lc code=end

