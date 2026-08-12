/*
 * @lc app=leetcode id=1103 lang=golang
 *
 * [1103] Distribute Candies to People
 */

// @lc code=start
func distributeCandies(candies int, num_people int) []int {
	i := 0
	j := 1
	res := make([]int, num_people)
	for candies > 0 {
		if candies < j {
			res[i] += candies
			break
		} else {
			res[i] += j
			i++
			candies -= j
			j++
			if i > num_people-1 {
				i = 0
			}
		}
	}
	return res
}

// @lc code=end

