/*
 * @lc app=leetcode id=941 lang=golang
 *
 * [941] Valid Mountain Array
 */

// @lc code=start
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	b := 0
	for i, v := range arr {
		if i == 0 {
			continue
		}
		if i == 1 {
			if v <= arr[0] {
				return false
			}
		} else if v < arr[i-1] && b == 0 {
			b = 1
		} else if v > arr[i-1] && b == 1 {
			return false
		} else if v == arr[i-1] {
			return false
		}
	}
	return b == 1
}

// @lc code=end

