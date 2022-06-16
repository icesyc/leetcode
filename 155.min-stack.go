/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
/**
 * 使用两个stack,一个正常stack,另一个保存当前stack的最小值，倒序
 * 如[6,9,4,8,2]
 * data=[6,9,4,8,2]
 * minData[6,4,2]
 * 当pop stack时，检查下是否是最小值，如果是的话就将其从minStack上弹出
 */
type MinStack struct {
	data    []int
	minData []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.minData) == 0 || val <= this.GetMin() {
		this.minData = append(this.minData, val)
	}
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	if this.Top() == this.GetMin() {
		this.minData = this.minData[:len(this.minData)-1]
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minData[len(this.minData)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

