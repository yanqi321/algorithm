package main

import "fmt"

func insertionSort(arr []int) []int {
	for i:=1; i < len(arr); i++ {
		prv := i - 1
		for prv >=0 && arr[prv] > arr[prv+1] {
			arr[prv], arr[prv+ 1] = arr[prv+1], arr[prv]
			prv--
		}
	}
	return arr
}

func main() {
	ret := insertionSort([]int{1, 3, 4, 5, 10, 100, 30, 40, 7, 11, 18})
	fmt.Println(ret, len(ret))
}
