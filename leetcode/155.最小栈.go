/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (52.34%)
 * Likes:    429
 * Dislikes: 0
 * Total Accepted:    88.4K
 * Total Submissions: 168.8K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
 * 
 * 
 * push(x) -- 将元素 x 推入栈中。
 * pop() -- 删除栈顶的元素。
 * top() -- 获取栈顶元素。
 * getMin() -- 检索栈中的最小元素。
 * 
 * 
 * 示例:
 * 
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 * 
 * 
 */

// @lc code=start
type MinStack struct {
  stack []Node
}

type Node struct {
  val int
  min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
  return MinStack{}
}


func (this *MinStack) Push(x int)  {
  node := Node{ x, x }
  if len(this.stack) > 0 && this.stack[len(this.stack) - 1].min < x {
    node.min = this.stack[len(this.stack) - 1].min
  }
  this.stack = append(this.stack, node)
}


func (this *MinStack) Pop()  {
  len := len(this.stack)
  this.stack = this.stack[:len - 1]
}


func (this *MinStack) Top() int {
  len := len(this.stack)
  return this.stack[len- 1].val
}


func (this *MinStack) GetMin() int {
  len := len(this.stack)
  return this.stack[len - 1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

