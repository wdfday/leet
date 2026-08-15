/*
 * @lc app=leetcode id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	inDegree := make(map[int]int)
	adjList := make(map[int][]int)

	for _, edge := range edges {
		inDegree[edge[0]]++
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		inDegree[edge[1]]++
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	leaves := []int{}
	for k, v := range inDegree {
		if v == 1 {
			leaves = append(leaves, k)
		}
	}

	remaining := n
	for remaining > 2 {
		remaining -= len(leaves)
		nextLeaves := []int{}
		for _, leaf := range leaves {
			for _, u := range adjList[leaf] {
				inDegree[u]--
				if inDegree[u] == 1 {
					nextLeaves = append(nextLeaves, u)
				}
			}
		}
		leaves = nextLeaves
	}

	return leaves
}
// @lc code=end

