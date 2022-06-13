/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 */
// @lc code=start
/**
 * 分配所需的最少糖果数
 * 根据题目可知，
 * 我们最开始分配的肯定是1, 如果ratings[i] > ratings[i-1]
 * 那么分配数量就要+1
 * 如果ratings[i] == ratings[i-1], 我们分配一个就行了
 * 如果ratings[i] < ratings[i-1]，我们分配1
 * 如果前面的已经分配了1，那么这时就要增加2个
 * 考虑以下例子[1,3,8,6,5,4]
 * 分配过程如下
 * 1
 * 1,2
 * 1,2,3 <-peak
 * 1,2,3,1
 * 1,2,3,2,1
 * 1,2,4,3,2,1
 * 定义三个变量
 * up 升序的次数
 * down 降序的次数
 * peak 升序的最后值
 * 三种情况分别计算应该分配的糖果数
 * 升序时需要分配 1+up个(因为up从0开始)
 * 降序时需要分配down个, 如果peak值小于down，peak对应的位置也需要+1，所以总数也需要+1
 * 平序时 重置up,down,peak 增加1即可
 */
/*
func candy(ratings []int) int {
	total := 1
	peak, up, down := 0, 0, 0
	for i := 1; i < len(ratings); i++ {
		//升序
		if ratings[i] > ratings[i-1] {
			up++
			down = 0
			total += 1 + up
			peak = up
		} else if ratings[i] == ratings[i-1] {
			peak, up, down = 0, 0, 0
			total += 1
		} else {
			up = 0
			down++
			total += down
			if peak < down {
				total += 1
			}
		}
	}
	return total
}
*/
/**
 * 解法2:
 * 初始化长度为n的数组，保存给每个人发的糖果数,初始为1
 * 先从左向右扫描, 保证右侧排名高的拿的糖果比左边多
 * 再从右往左扫描，保证左侧排名高的拿的糖果比右侧多
 * 最后把数组加和
 */
func candy(ratings []int) int {
	n := len(ratings)
	nums := make([]int, n)
	nums[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			nums[i] = nums[i-1] + 1
		} else {
			nums[i] = 1
		}
	}
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			if nums[i]+1 > nums[i-1] {
				nums[i-1] = nums[i] + 1
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res += nums[i]
	}
	return res
}

// @lc code=end

