/*
 * @lc app=leetcode id=402 lang=golang
 *
 * [402] Remove K Digits
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	stack := make([]byte, 0, len(num))

	for i := 0; i < len(num); i++ {
		d := num[i]

		for len(stack) > 0 &&
			k > 0 &&
			stack[len(stack)-1] > d {

			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, d)
	}

	// nếu vẫn còn quota xóa
	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	// bỏ leading zero
	idx := 0
	for idx < len(stack) && stack[idx] == '0' {
		idx++
	}

	if idx == len(stack) {
		return "0"
	}

	return string(stack[idx:])
}

// @lc code=end
