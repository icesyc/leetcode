/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
/**
 * 每个课程及其依赖的课程，可以组成一个图
 * 如果有互相依赖，则不能完成课程
 * 所以答案为判断该图是否有环
 * 图采用邻接数组表示, 索引为前置课程，值为后续课程
 * graph[0] = [1,2]
 * graph[1] = [3,4]
 * graph[2] = [5]
 * 使用dfs遍历图，为了防止重复使用visited保存已经遍历过的节点
 * 使用path保存每次traverse时经过的节点，如果path中存在说明有环
 */
/*
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	visited := make([]bool, numCourses)
	for i := 0; i < len(graph); i++ {
		path := make([]bool, numCourses)
		hasCycle := !traverse(graph, i, visited, path)
		if hasCycle {
			return false
		}
	}
	return true
}

func traverse(graph [][]int, node int, visited []bool, path []bool) bool {
	if path[node] {
		return false
	}
	if visited[node] {
		return true
	}
	visited[node] = true
	path[node] = true
	for _, n := range graph[node] {
		if !traverse(graph, n, visited, path) {
			return false
		}
	}
	path[node] = false
	return true
}

//将课程条件转换成图
func buildGraph(numCosuses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCosuses)
	for _, req := range prerequisites {
		//from为先要完成的课程
		from, to := req[1], req[0]
		graph[from] = append(graph[from], to)
	}
	return graph
}
*/
/**
 * 方法2 判断有向图是否可以拓扑排序
 * 拓扑排序的结果是方向后面的节点一定排在前面的节点后面
 * 首先计算各个节点的入度(有多少节点指向它)
 * 拓扑排序是先将所有indegree为0的节点加入到queue中
 * 从queue中取出节点
 * 将该节点连接的所有节点入度-1
 * 如果连接的节点入度为0，那么加入到queue中
 * 如果遍历完成后的count数不等于图的结点数，那么图不能使用拓扑排序
 */
func canFinish(numCourses int, prerequisites [][]int) bool {
	//构建图，同时计算degree
	graph := make([][]int, numCourses)
	degree := make([]int, numCourses)
	for _, req := range prerequisites {
		from, to := req[1], req[0]
		graph[from] = append(graph[from], to)
		degree[to]++
	}

	//找出degree为0的节点
	queue := []int{}
	for i := 0; i < len(degree); i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, edge := range graph[node] {
			degree[edge]--
			if degree[edge] == 0 {
				queue = append(queue, edge)
			}
		}
		count++
	}
	return count == numCourses
}

// @lc code=end

