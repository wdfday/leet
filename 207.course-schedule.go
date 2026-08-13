/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {

	adj := make([][]int, numCourses)
	inEdge := make(map[int]int)
	for _, prerequisite := range prerequisites {
		adj[prerequisite[1]] = append(adj[prerequisite[1]], prerequisite[0])
		inEdge[prerequisite[0]]++
	}

	q := []int{}

	for i := range numCourses {
		if inEdge[i] == 0 {
			q = append(q, i)
		}
	}
    idx := 0
	for len(q) > idx {
		cur := q[idx]
		for _, v := range adj[cur] {
			inEdge[v]--
			if inEdge[v] == 0 {
				q = append(q, v)
			}
		}
		idx++
	}

	if len(q) == numCourses {
		return true
	}
	return false
}
// @lc code=end

