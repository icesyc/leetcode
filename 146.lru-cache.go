/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
/**
 * LRUcache的典型数据结构是使用map+双向链表
 * map存储链表结点, 保证get的时间复杂度为O(1)
 * 双向链表用来实现LRU, 每次get将key移到队头
 * put的时候也是put到队头，如果容量超了，从队尾移出一个key
 */
type LRUCache struct {
	Cache map[int]*Node
	Head  *Node
	Tail  *Node
	Cap   int
}
type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Cache: make(map[int]*Node),
		Cap:   capacity,
	}
	cache.Head, cache.Tail = &Node{}, &Node{}
	cache.Head.Next = cache.Tail
	cache.Tail.Prev = cache.Head
	return cache
}

//添加到队头
func (this *LRUCache) AddNode(node *Node) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next.Prev = node
	this.Head.Next = node
}

//移动到队头
func (this *LRUCache) MoveToHead(node *Node) {
	this.RemoveNode(node)
	this.AddNode(node)
}

//删除节点
func (this *LRUCache) RemoveNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
func (this *LRUCache) PopTail() *Node {
	node := this.Tail.Prev
	this.RemoveNode(node)
	return node
}

//get时需要把key移到队头
func (this *LRUCache) Get(key int) int {
	node := this.Cache[key]
	if node == nil {
		return -1
	}
	this.MoveToHead(node)
	return node.Val
}

//put时添加到队头,如果超过cap需要从队尾删除一个key
func (this *LRUCache) Put(key int, value int) {
	node := this.Cache[key]
	if node != nil {
		node.Val = value
		this.MoveToHead(node)
		return
	}
	node = &Node{Key: key, Val: value}
	this.Cache[key] = node
	this.AddNode(node)
	if this.Cap < len(this.Cache) {
		tail := this.PopTail()
		delete(this.Cache, tail.Key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

