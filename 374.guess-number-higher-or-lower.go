/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {

	l, r := 1, int(1<<31-1)

	for l < r {
		mid := l + (r-l)/2
		g := guess(mid)
		if g == 0 {
			return mid
		} else if g == 1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

// @lc code=end

