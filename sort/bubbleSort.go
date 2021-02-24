package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := range arr {
		didSwap := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				didSwap = true
			}
		}
		if !didSwap { // 优化 最好复杂度为 n
			return arr
		}
	}
	return arr
}

func main() {
	ret := bubbleSort([]int{1, 3, 4, 5, 10, 100, 30, 40, 7, 11, 18})
	fmt.Println(ret, len(ret))
}
