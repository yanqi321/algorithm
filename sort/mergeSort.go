package main

import "fmt"

func mergeSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen < 2 {
		return arr
	}
	m := arrLen / 2
	return merge(mergeSort(arr[:m]), mergeSort(arr[m:]))
}

func merge(l, r []int) []int {
	ret := make([]int, 0)
	for len(l) != 0 && len(r) != 0 {
		if l[0] < r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	if len(l) != 0 {
		ret = append(ret, l...)
	} else {
		ret = append(ret, r...)
	}
	fmt.Println("level",ret)
	return ret
}

func main() {
	ret := mergeSort([]int{1,3,4,5,10,100,30,40,7,11,18})
	fmt.Println(ret, len(ret))
}