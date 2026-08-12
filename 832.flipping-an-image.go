/*
 * @lc app=leetcode id=832 lang=golang
 *
 * [832] Flipping an Image
 */

// @lc code=start
func flipAndInvertImage(image [][]int) [][]int {
	for _, mage := range image {
		l, r := 0, len(mage)-1
		for l <= r {
			mage[l], mage[r] = 1-mage[r], 1-mage[l]
			l++
			r--
		}
	}
	return image
}

// @lc code=end

