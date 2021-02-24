package main

import "fmt"
// 希尔排序原理相关文档：https://www.cnblogs.com/chengxiao/p/6104371.html 不稳定排序
func shellSort(arr []int) []int {
	N := len(arr)
	for gap := N / 2; gap > 0; gap /=2 {
		for i := gap; i < N; i++ {
			j := i
			for j - gap >= 0 && arr[j] < arr[j - gap] { // 插入排序 从后往前 交换位置
				arr[j], arr[j - gap] = arr[j - gap], arr[j]
				j -= gap
			}
			// current := arr[j]
			// for j - gap >= 0 && current < arr[j - gap] { // 插入排序 从后往前 移动位置
			// 	arr[j] = arr[j - gap]
			// 	j -= gap
			// }
			// arr[j] = current
		}
	}
	return arr
}

func main() {
	ret := shellSort([]int{1, 3, 4, 5, 10, 100, 30, 40, 7, 11, 18})
	fmt.Println(ret, len(ret))
}
