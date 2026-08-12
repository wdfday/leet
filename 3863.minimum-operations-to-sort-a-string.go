/*
 * @lc app=leetcode id=3863 lang=golang
 *
 * [3863] Minimum Operations to Sort a String
 */

// @lc code=start
func minOperations(s string) int {
	n := len(s)

	// Bước 1: đã sorted chưa
	sorted := true
	for i := 1; i < n; i++ {
		if s[i] < s[i-1] {
			sorted = false
			break
		}
	}
	if sorted {
		return 0
	}

	// Bước 2: n == 2 và chưa sorted -> bất khả thi
	if n == 2 {
		return -1
	}

	// Tìm mn, mx và vị trí xuất hiện đầu/cuối của chúng
	mn, mx := s[0], s[0]
	for i := 1; i < n; i++ {
		if s[i] < mn {
			mn = s[i]
		}
		if s[i] > mx {
			mx = s[i]
		}
	}

	// Bước 3: 1 thao tác
	if s[0] == mn || s[n-1] == mx {
		return 1
	}

	// Bước 4: 2 thao tác — mn hoặc mx nằm ở giữa
	hasMidMn, hasMidMx := false, false
	for i := 1; i < n-1; i++ {
		if s[i] == mn {
			hasMidMn = true
		}
		if s[i] == mx {
			hasMidMx = true
		}
	}
	if hasMidMn || hasMidMx {
		return 2
	}

	// Bước 5: còn lại -> mx ở đầu, mn ở cuối -> 3 thao tác
	return 3
}

// @lc code=end

