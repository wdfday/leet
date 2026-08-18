/*
 * @lc app=leetcode id=354 lang=golang
 *
 * [354] Russian Doll Envelopes
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
    sort.Slice(envelopes, func(i, j int) bool {
        if envelopes[i][0] == envelopes[j][0] {
            return envelopes[i][1] > envelopes[j][1] // giảm dần
        }
        return envelopes[i][0] < envelopes[j][0]
    })

    tails := []int{}

    for _, e := range envelopes {
        h := e[1]

        idx := sort.SearchInts(tails, h)

        if idx == len(tails) {
            tails = append(tails, h)
        } else {
            tails[idx] = h
        }
    }

    return len(tails)

}
// @lc code=end

