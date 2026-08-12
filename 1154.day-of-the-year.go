/*
 * @lc app=leetcode id=1154 lang=golang
 *
 * [1154] Day of the Year
 */

// @lc code=start
func dayOfYear(date string) int {
	year := 0
	month := 0
	day := 0

	fmt.Sscanf(date, "%d-%d-%d", &year, &month, &day)

	days := [...]int{
		31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31,
	}

	if month > 2 && isLeap(year) {
		days[1]++
	}

	for i := 0; i < month-1; i++ {
		day += days[i]
	}

	return day
}

func isLeap(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
// @lc code=end

