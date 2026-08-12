/*
 * @lc app=leetcode id=1071 lang=golang
 *
 * [1071] Greatest Common Divisor of Strings
 */

// @lc code=start
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	// Điều kiện cần và đủ: str1 và str2 phải "ghép hoán vị" ra kết quả giống nhau
	if str1+str2 != str2+str1 {
		return ""
	}
	// Độ dài chuỗi GCD = gcd(len1, len2)
	g := gcd(len(str1), len(str2))
	return str1[:g]
}
// @lc code=end

