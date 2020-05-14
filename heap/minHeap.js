class MinHeap {
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
    if (l < this.len && this.more(maxChildIndex, l)) {
      maxChildIndex = l
    }
    if (r < this.len && this.more(maxChildIndex, r)) {
      maxChildIndex = r
    }
    if (maxChildIndex != index) {
      this.swap(index, maxChildIndex)
      this.shiftDown(maxChildIndex)
    }
  }
  shiftUp(index) {
    const parent = parseInt(index/2 -1) // parent 可能为 -0.5 不能用floor
    if ( index > 0 && this.more(parent, index)) {
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

var findKthLargest = function (nums, k) {
  let minHeap = new MinHeap()

  for (let i = 0; i < nums.length; i++) {
    if (minHeap.size() < k) {
      minHeap.push(nums[i])
    } else if (minHeap.top() < nums[i]) {
      minHeap.pop()
      minHeap.push(nums[i])
    }
  }
  return minHeap.top()
};

const val = findKthLargest([1,3,5,7,9,2,11,22,55,99], 6)
console.log(val)