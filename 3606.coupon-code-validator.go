/*
 * @lc app=leetcode id=3606 lang=golang
 *
 * [3606] Coupon Code Validator
 */

// @lc code=start
func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	priority := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}

	isValidCode := func(s string) bool {
		if len(s) == 0 {
			return false
		}
		for _, c := range s {
			if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
				return false
			}
		}
		return true
	}

	type coupon struct {
		code string
		rank int
	}

	var valid []coupon
	for i, c := range code {
		rank, ok := priority[businessLine[i]]
		if !ok || !isActive[i] || !isValidCode(c) {
			continue
		}
		valid = append(valid, coupon{code: c, rank: rank})
	}

	sort.Slice(valid, func(i, j int) bool {
		if valid[i].rank != valid[j].rank {
			return valid[i].rank < valid[j].rank
		}
		return valid[i].code < valid[j].code
	})

	res := make([]string, len(valid))
	for i, v := range valid {
		res[i] = v.code
	}
	return res
}
// @lc code=end

