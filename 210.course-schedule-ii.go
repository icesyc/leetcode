/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start
/**
 * 同207题，就是拓扑排序的结果
 * 也可以使用dfs遍历，后序处理结果
 */
func findOrder(numCourses int, prerequisites [][]int) []int {
	//建图、生成入度数据
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, req := range prerequisites {
		from, to := req[1], req[0]
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	//先入队indegree为0的节点
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	res := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)
		for _, edge := range graph[node] {
			indegree[edge]--
			if indegree[edge] == 0 {
				queue = append(queue, edge)
			}
		}
	}
	if len(res) == numCourses {
		return res
	}
	return []int{}
}

// @lc code=end

