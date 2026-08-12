/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	prev := 0 // giá trị ô liền trước, coi biên trái như đã có sẵn 0

	for i := 0; i < len(flowerbed); i++ {
		next := 0
		if i < len(flowerbed)-1 {
			next = flowerbed[i+1]
		}
		if flowerbed[i] == 0 && prev == 0 && next == 0 {
			flowerbed[i] = 1
			count++
			if count >= n {
				return true
			}
			prev = 1
		} else {
			prev = flowerbed[i]
		}
	}

	return count >= n
}

// @lc code=end

