package main

import "container/heap"

type IntMinHeap []int
type KthLargest struct {
	k int
	heap IntMinHeap
}
func (hp IntMinHeap) Len() int {
	return len(hp)
}
func (hp IntMinHeap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}
func (hp *IntMinHeap) Push(x interface{}) {
	*hp = append(*hp, x.(int))
}
func (hp *IntMinHeap) Pop() interface{} {
	tmp := *hp
	n := len(tmp)
	x := (tmp)[n-1]
	*hp = (tmp)[0:n-1]
	return x
}
func (hp IntMinHeap) Less(i, j int) bool {
	return	hp[i] < hp[j]
}

func Constructor(k int, nums []int) KthLargest{
	var mp IntMinHeap
	for i, n := range nums {
		if i <= k - 1 {
			heap.Push(&mp, n)
		} else if (n < mp[0]) {
			heap.Pop(&mp)
			heap.Push(&mp, n)
		}
	}
	return KthLargest{k, mp}
}

func (kth *KthLargest) Add(val int) int {
	if (len(kth.heap) < kth.k) {
		heap.Push(&kth.heap, val)
	} else if (val < kth.heap[0]) {
		heap.Pop(&kth.heap)
		heap.Push(&kth.heap, val)
	}
	return kth.heap[0]
}