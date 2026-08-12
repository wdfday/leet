/*
 * @lc app=leetcode id=929 lang=golang
 *
 * [929] Unique Email Addresses
 */

// @lc code=start
func numUniqueEmails(emails []string) int {
	m := make(map[string]struct{})
	for _, email := range emails {
		at := strings.IndexByte(email, '@')
		local, domain := email[:at], email[at:]

		if plus := strings.IndexByte(local, '+'); plus != -1 {
			local = local[:plus]
		}
		local = strings.ReplaceAll(local, ".", "")

		m[local+domain] = struct{}{}
	}
	return len(m)
}
// @lc code=end

