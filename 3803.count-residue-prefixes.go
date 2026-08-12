/*
 * @lc app=leetcode id=3803 lang=golang
 *
 * [3803] Count Residue Prefixes
 */

// @lc code=start
func residuePrefixes(s string) int {
    seen := make(map[byte]bool)
    distinct := 0
    ans := 0

    for i := 0; i < len(s); i++ {
        if !seen[s[i]] {
            seen[s[i]] = true
            distinct++
        }

        length := i + 1

        if distinct == length%3 {
            ans++
        }
    }

    return ans

}
// @lc code=end

