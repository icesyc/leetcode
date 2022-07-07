/*
 * @lc app=leetcode id=218 lang=golang
 *
 * [218] The Skyline Problem
 */

// @lc code=start
/**
 * 基本思路是排序+优先队列
 * 先将所有build的坐标转换为[x, height]的结构对
 * [left, right, height] => [left, height], [right,height]
 * 对所有坐标点排序，排序的规则如下
 * 按x坐标从小到大排序
 * 对于x相同的，按height倒排
 * 为了使边界case排序能够正确，我们在插入的时候对于结束点取[right,-height]
 *
 * 然后遍历所有points, 每次遇到起点，将height添加到队列中,遇到终点从队列中移除该点对应的height
 * 添加或删除操作后，如果队列的最大值发生变化，说明需要记录该坐标
 *
 * 边界case有三种， 假设有A,B两个建筑物
 * 1. 如果A, B的起点相同，A.height < B.height, 那么key point应该为B.height
 * 在添加到优先队列时，应该先添加B,如果先添加A，高度同样发生变化，A.height也会成为key point，结果是错的
 * 2. 如果A,B的终点相同, A.height < B.height
 * 那么从队列移除时，应该先移除A.height, 如果先移除B, 那么高度发生变化，B也会被添加到结果中
 * 3. 如果A的终点和B的起点相同，A.height < B.height
 * 那么队列中的顺序应该为先添加B，再删除A, 删除A的时候队列的最大值不会变化，不会将A做为key point
 *
 * 总结一下就是添加的时候，先添加height大的，删除的时候先删除height小的
 * 这样可以保证小的height入队出队时不会成为最大值，也就不会被当做key point处理
 * 为了达到这个效果，我们在转换坐标点时，对于起点取height, 终点取-height
 * 起点相同height大的排在前面
 * 终点相同height小的排在前面
 * 起终相同起点排在前面
 *
 * 参考视频https://youtu.be/GSBLe8cKu0s?t=621
 */
//*
type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	x := old[len(old)-1]
	*pq = old[:len(old)-1]
	return x
}
func (pq *PriorityQueue) Add(x int) {
	heap.Push(pq, x)
}
func (pq *PriorityQueue) Remove(x int) {
	for i := 0; i < len(*pq); i++ {
		if (*pq)[i] == x {
			heap.Remove(pq, i)
			return
		}
	}
}
func (pq *PriorityQueue) Peak() int {
	return (*pq)[0]
}
func getSkyline(buildings [][]int) [][]int {
	res := [][]int{}
	points := [][]int{}
	//转换成points
	for i := 0; i < len(buildings); i++ {
		points = append(points, []int{buildings[i][0], buildings[i][2]})
		points = append(points, []int{buildings[i][1], -buildings[i][2]})
	}
	//排序，先按x坐标，再按height倒序
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	pq := &PriorityQueue{0}
	preHeight := 0
	for _, p := range points {
		//起点，添加到queue中
		if p[1] > 0 {
			pq.Add(p[1])
		} else {
			//终点，从queue中移除起点的height
			pq.Remove(-p[1])
		}
		//最大高度有变化，添加key point
		if pq.Peak() != preHeight {
			preHeight = pq.Peak()
			res = append(res, []int{p[0], pq.Peak()})
		}
	}
	return res
}

//*/
/**
 * 解法2
 * 将所有building的left,right坐标保存到数组中，过滤重复元素并排序
 * 遍历数组所有point, 找出当前位置的最高位置(遍历builds, build.left <=x && build.right > x)
 * 如果最高位置发生变化，p为key point
 * 参考
 * https://leetcode.com/problems/the-skyline-problem/discuss/395923/JavaScript-Easy-and-Straightforward-with-picture-illustrations
 */
/*
func getSkyline(buildings [][]int) [][]int {
	set := map[int]bool{}
	for _, b := range buildings {
		set[b[0]] = true
		set[b[1]] = true
	}
	points := []int{}
	for k := range set {
		points = append(points, k)
	}
	sort.Ints(points)
	res := [][]int{}
	preHeight := 0
	for _, p := range points {
		height := 0
		for i := 0; i < len(buildings) && buildings[i][0] <= p; i++ {
			if buildings[i][1] > p {
				height = max(height, buildings[i][2])
			}
		}
		if height != preHeight {
			res = append(res, []int{p, height})
			preHeight = height
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

// @lc code=end

