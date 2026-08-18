/*
 * @lc app=leetcode id=393 lang=golang
 *
 * [393] UTF-8 Validation
 */

// @lc code=start
func validUtf8(data []int) bool {
    e7 := 1 << 7
    e6 := 1 << 6
    e5 := 1 << 5
    e4 := 1 << 4
	e3 := 1 << 3 

    for i := 0; i < len(data); i++ {
        b := data[i]

        if b < e7 {
            continue
        }

        tail := 0

        if b&e6 != 0 {
            tail++
            if b&e5 != 0 {
                tail++
                if b&e4 != 0 {
                    tail++
					if b&e3 != 0 {
						return false // 11111xxx
					}
                }
            }
        }

        if tail == 0 || tail > 3 {
            return false
        }

        for j := 1; j <= tail; j++ {
            if i+j >= len(data) || data[i+j]&e7 == 0 || data[i+j]&e6 != 0 {
                return false
            }
        }

        i += tail
    }

    return true
}
// @lc code=end

