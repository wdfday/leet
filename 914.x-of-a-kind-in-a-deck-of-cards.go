/*
 * @lc app=leetcode id=914 lang=golang
 *
 * [914] X of a Kind in a Deck of Cards
 */

// @lc code=start
func hasGroupsSizeX(deck []int) bool {
    freq := make(map[int]int)

    for _, x := range deck {
        freq[x]++
    }

    g := 0
    for _, count := range freq {
        g = gcd(g, count)
    }

    return g > 1
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}
// @lc code=end

