/*
 * @lc app=leetcode id=3005 lang=golang
 *
 * [3005] Count Elements With Maximum Frequency
 */

// @lc code=start
func maxFrequencyElements(nums []int) int {
	
	res := 0
	maxF := 0
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
		if freq[num] > maxF {
			res = freq[num]
			maxF = freq[num]
		} else if freq[num] == maxF {
			res += maxF
		}
	}
	return res

    
}
// @lc code=end

