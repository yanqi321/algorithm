/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 *
 * https://leetcode-cn.com/problems/design-circular-deque/description/
 *
 * algorithms
 * Medium (51.18%)
 * Likes:    27
 * Dislikes: 0
 * Total Accepted:    6.3K
 * Total Submissions: 12.4K
 * Testcase Example:  '["MyCircularDeque","insertLast","insertLast","insertFront","insertFront","getRear","isFull","deleteLast","insertFront","getFront"]\n' +
  '[[3],[1],[2],[3],[4],[],[],[],[4],[]]'
 *
 * 设计实现双端队列。
 * 你的实现需要支持以下操作：
 * 
 * 
 * MyCircularDeque(k)：构造函数,双端队列的大小为k。
 * insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
 * insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
 * deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
 * deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
 * getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
 * getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
 * isEmpty()：检查双端队列是否为空。
 * isFull()：检查双端队列是否满了。
 * 
 * 
 * 示例：
 * 
 * MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
 * circularDeque.insertLast(1);                    // 返回 true
 * circularDeque.insertLast(2);                    // 返回 true
 * circularDeque.insertFront(3);                    // 返回 true
 * circularDeque.insertFront(4);                    // 已经满了，返回 false
 * circularDeque.getRear();                  // 返回 2
 * circularDeque.isFull();                        // 返回 true
 * circularDeque.deleteLast();                    // 返回 true
 * circularDeque.insertFront(4);                    // 返回 true
 * circularDeque.getFront();                // 返回 4
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 所有值的范围为 [1, 1000]
 * 操作次数的范围为 [1, 1000]
 * 请不要使用内置的双端队列库。
 * 
 * 
 */

// @lc code=start
type MyCircularDeque struct {
  head *Node
  rear *Node
  len int
  count int
}

type Node struct {
  Prev *Node
  Next *Node
  Val int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
  deque := MyCircularDeque{
    head: nil,
    rear: nil,
    len: k,
    count: 0,
  }
  return deque
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
  if this.IsFull() {
    return false
  }
  node := Node{
    Prev: nil,
    Next: this.head,
    Val: value,
  }
  if this.IsEmpty() {
    this.rear = &node
  } else  {
    this.head.Prev = &node
  }
  this.head = &node
  this.count++
  return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
  if this.IsFull() {
    return false
  }
  node := Node{
    Prev: this.rear,
    Next: nil,
    Val: value,
  }
  if this.IsEmpty() {
    this.head = &node
  } else {
    this.rear.Next = &node
  }
  this.rear = &node
  this.count++
  return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
  if this.IsEmpty() {
    return false
  }
  this.head = this.head.Next
  if this.head == nil {
    this.rear = nil
  } else {
    this.head.Prev = nil
  }
  this.count--
  return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
  if this.IsEmpty() {
    return false
  }
  this.rear = this.rear.Prev
  if this.rear == nil {
    this.head = nil
  } else {
      this.rear.Next = nil
  }
  this.count--
  return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
  if this.IsEmpty() {
    return -1
  }
  return this.head.Val
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
  if this.IsEmpty() {
    return -1
  }
  return this.rear.Val
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
  return this.count == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
  return this.len == this.count
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

