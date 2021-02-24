package main

import "fmt"

func selectionSort(arr []int) []int {
	for i,v := range arr {
		for j:=i; j < len(arr);j++ {
			if arr[j] < v {
				arr[j],arr[i] = v,arr[j]
				v = arr[i]
			}
		}
	}
	return arr
}

func main() {
	ret := selectionSort([]int{1,3,4,5,10,100,30,40,7,11,18})
	fmt.Println(ret, len(ret))
}