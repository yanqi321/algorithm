class MaxHeap {
  constructor() {
    this.heap = []
    this.len = 0
  }
  push(val) {
    const index = ++this.len - 1
    this.heap[index] = val
    this.shiftUp(index)
  }
  pop() {
    const top = this.heap[0]
    this.swap(0, --this.len)
    this.heap[this.len] = undefined
    this.shiftDown(0)
    return top
  }
  shiftDown(index) {
    const l = 2 * index + 1
    const r = 2 * index + 2
    let maxChildIndex = index
    if (l < this.len && this.more(l, maxChildIndex)) {
      maxChildIndex = l
    }
    if (r < this.len && this.more(r, maxChildIndex)) {
      maxChildIndex = r
    }
    if (maxChildIndex != index) {
      this.swap(index, maxChildIndex)
      this.shiftDown(maxChildIndex)
    }
  }
  shiftUp(index) {
    const parent = parseInt(index/2 -1) // parent 可能为 -0.5 不能用floor
    if ( index > 0 && this.more(index, parent)) {
      this.swap(parent, index)
      this.shiftUp(parent)
    }
  }
  swap(i, j) {
    let temp = this.heap[i]
    this.heap[i] = this.heap[j]
    this.heap[j] = temp
  }
  more(i, j) {
    return this.heap[i] > this.heap[j]
  }
  top() {
    return this.heap[0]
  }
  isEmpty() {
    return this.len === 0
  }
  size() {
    return this.len
  }
}

var findKthLeast = function (nums, k) {
  let maxHeap = new MaxHeap()

  for (let i = 0; i < nums.length; i++) {
    if (maxHeap.size() < k) {
      maxHeap.push(nums[i])
    } else if (maxHeap.top() > nums[i]) {
      maxHeap.pop()
      maxHeap.push(nums[i])
    }
  }
  return maxHeap.top()
};

const val = findKthLeast([1,3,5,7,9,2,11,22,55,99], 5)
console.log(val)