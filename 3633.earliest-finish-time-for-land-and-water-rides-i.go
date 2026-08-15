/*
 * @lc app=leetcode id=3633 lang=golang
 *
 * [3633] Earliest Finish Time for Land and Water Rides I
 */

// @lc code=start
func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    ans := math.MaxInt

	for i := 0; i < len(landStartTime); i++ {
		for j := 0; j < len(waterStartTime); j++ {

			// Land -> Water
			landFinish := landStartTime[i] + landDuration[i]
			waterFinish := max(landFinish, waterStartTime[j]) + waterDuration[j]

			ans = min(ans, waterFinish)

			// Water -> Land
			waterFinish = waterStartTime[j] + waterDuration[j]
			landFinish = max(waterFinish, landStartTime[i]) + landDuration[i]

			ans = min(ans, landFinish)
		}
	}

	return ans
}
// @lc code=end

