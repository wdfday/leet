/*
 * @lc app=leetcode id=682 lang=golang
 *
 * [682] Baseball Game
 */

// @lc code=start
func calPoints(operations []string) int {
	a := make([]int, 0)
	for _, operation := range operations {
		switch operation {
		case "C":
			if len(a) > 0 {
				a = a[:len(a)-1]
			}
		case "D":
			if len(a) > 0 {
				a = append(a, a[len(a)-1]*2)
			}
		case "+":
			if len(a) > 1 {
				a = append(a, a[len(a)-1]+a[len(a)-2])
			}
		default:
			v, ok := strconv.Atoi(operation)
			if ok == nil {
				a = append(a, v)
			}
		}
	}

	res := 0
	for _, v := range a {
		res += v
	}
	return res
}

// @lc code=end

