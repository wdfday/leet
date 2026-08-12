/*
 * @lc app=leetcode id=1108 lang=golang
 *
 * [1108] Defanging an IP Address
 */

// @lc code=start
func defangIPaddr(address string) string {
    return strings.ReplaceAll(address, ".", "[.]")
}
// @lc code=end

