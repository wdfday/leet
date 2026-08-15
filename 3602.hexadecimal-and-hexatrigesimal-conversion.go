/*
 * @lc app=leetcode id=3602 lang=golang
 *
 * [3602] Hexadecimal and Hexatrigesimal Conversion
 */

// @lc code=start
func concatHex36(n int) string {
	hex := n * n
	het := hex * n 
	sHex := []byte{}
	sHet := []byte{}

	for hex > 0 {
		c := hex % 16
		var b byte
		if c < 10 {
			b = byte(c + '0')
		} else {
			b = byte(c - 10) + 'A'
		}
		sHex = append([]byte{b}, sHex...)
		hex /= 16
	}

	for het > 0 {
		c := het % 36
		var b byte
		if c < 10 {
			b = byte(c + '0')
		} else {
			b = byte(c - 10) + 'A'
		}
		sHet = append([]byte{b}, sHet...)
		het /= 36
	}

	return string(sHex) + string(sHet)
    
}
// @lc code=end

