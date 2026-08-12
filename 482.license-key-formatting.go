/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 */

// @lc code=start
func licenseKeyFormatting(s string, k int) string {
    s = strings.ReplaceAll(s, "-", "")
    s = strings.ToUpper(s)

    n := len(s)
    firstGroupLen := n % k

    var sb strings.Builder
    if firstGroupLen > 0 {
        sb.WriteString(s[:firstGroupLen])
    }
    for i := firstGroupLen; i < n; i += k {
        if sb.Len() > 0 {
            sb.WriteByte('-')
        }
        sb.WriteString(s[i : i+k])
    }
    return sb.String()
}

// @lc code=end

