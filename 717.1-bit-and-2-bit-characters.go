/*
 * @lc app=leetcode id=717 lang=golang
 *
 * [717] 1-bit and 2-bit Characters
 */

// @lc code=start
func isOneBitCharacter(bits []int) bool {

	for i :=0; i < len(bits); i++ {
		if bits[i] == 0 {
			continue
		} else {
			if i == len(bits) - 2 {
				return false
			}
			i++
		}
	}
    return true
}
// @lc code=end

