/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
/**
 * bfs搜索，使用map保存访问过的节点
 * 如果未访问过，就copy一个，并将其加入到queue中
 */
/*
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := []*Node{node}
	visited := map[int]*Node{}
	copy := &Node{Val: node.Val}
	visited[copy.Val] = copy
	//bfs搜索
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		clone := visited[cur.Val]
		for _, n := range cur.Neighbors {
			nb := visited[n.Val]
			if nb == nil {
				nb = &Node{Val: n.Val}
				visited[nb.Val] = nb
				queue = append(queue, n)
			}
			clone.Neighbors = append(clone.Neighbors, nb)
		}
	}
	return copy
}
*/

//dfs搜索，使用递归
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := map[int]*Node{}
	visited[node.Val] = &Node{Val: node.Val}
	return dfs(node, visited)
}

func dfs(node *Node, visited map[int]*Node) *Node {
	copy := visited[node.Val]
	for _, n := range node.Neighbors {
		nb := visited[n.Val]
		if nb == nil {
			nb = &Node{Val: n.Val}
			visited[nb.Val] = nb
			dfs(n, visited)
		}
		copy.Neighbors = append(copy.Neighbors, nb)
	}
	return copy
}

// @lc code=end

