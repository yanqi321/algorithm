/*
 * @lc app=leetcode.cn id=480 lang=golang
 *
 * [480] 滑动窗口中位数
 */

// @lc code=start
func medianSlidingWindow(nums []int, k int) []float64 {
	ans := []float64{}
	minH := MinHeapR{}
	maxH := MaxHeapR{}
	for i := range nums {
		if minH.Len() == 0 || nums[i] > minH.Top() {
			minH.Push(nums[i])
		} else {
			maxH.Push(nums[i])
		}
		if i >= k {
			if nums[i - k] >= minH.Top() {
				minH.Remove(nums[i - k])
			} else {
				maxH.Remove(nums[i - k])
			}
		}
		if minH.Len() > maxH.Len() + 1 {
			maxH.Push(minH.Pop())
		} else if minH.Len() < maxH.Len() {
			minH.Push(maxH.Pop())
		}
		if minH.Len() + maxH.Len() == k {
			if k & 1 == 0 {
				ans = append(ans, float64(minH.Top() + maxH.Top()) / 2.0)
			} else {
				ans = append(ans, float64(minH.Top()))
			}
		}
	}
	return ans
}

// IntHeap define
type IntHeap struct {
	data []int
}

// Len define
func (h IntHeap) Len() int { return len(h.data) }

// Swap define
func (h IntHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

// Push define
func (h *IntHeap) Push(x interface{}) { h.data = append(h.data, x.(int)) }

// Pop define
func (h *IntHeap) Pop() interface{} {
	x := h.data[h.Len()-1]
	h.data = h.data[0 : h.Len()-1]
	return x
}

// Top defines
func (h IntHeap) Top() int {
	return h.data[0]
}

// MinHeap define
type MinHeap struct {
	IntHeap
}

// Less define
func (h MinHeap) Less(i, j int) bool { return h.data[i] < h.data[j] }

// MaxHeap define
type MaxHeap struct {
	IntHeap
}

// Less define
func (h MaxHeap) Less(i, j int) bool { return h.data[i] > h.data[j] }

// MinHeapR define
type MinHeapR struct {
	hp, hpDel MinHeap
}

// Len define
func (h MinHeapR) Len() int { return h.hp.Len() - h.hpDel.Len() }

// Top define
func (h *MinHeapR) Top() int {
	for h.hpDel.Len() > 0 && h.hp.Top() == h.hpDel.Top() {
		heap.Pop(&h.hp)
		heap.Pop(&h.hpDel)
	}
	return h.hp.Top()
}

// Pop define
func (h *MinHeapR) Pop() int {
	x := h.Top()
	heap.Pop(&h.hp)
	return x
}

// Push define
func (h *MinHeapR) Push(x int) { heap.Push(&h.hp, x) }

// Remove define
func (h *MinHeapR) Remove(x int) { heap.Push(&h.hpDel, x) }

// MaxHeapR define
type MaxHeapR struct {
	hp, hpDel MaxHeap
}

// Len define
func (h MaxHeapR) Len() int { return h.hp.Len() - h.hpDel.Len() }

// Top define
func (h *MaxHeapR) Top() int {
	for h.hpDel.Len() > 0 && h.hp.Top() == h.hpDel.Top() {
		heap.Pop(&h.hp)
		heap.Pop(&h.hpDel)
	}
	return h.hp.Top()
}

// Pop define
func (h *MaxHeapR) Pop() int {
	x := h.Top()
	heap.Pop(&h.hp)
	return x
}

// Push define
func (h *MaxHeapR) Push(x int) { heap.Push(&h.hp, x) }

// Remove define
func (h *MaxHeapR) Remove(x int) { heap.Push(&h.hpDel, x) }

// @lc code=end

