package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//arr := []int{4, 5, 7, 2, 3, 57, 8, 1, 34, 6, 756, 1, 3, 358, 4, 67, 58, 3, 1, 7, 3, 34, 3, 6, 6}
	arr:=randomArr()
	fmt.Println("排序前：", arr)

	tStart := time.Now()

	QuickSort(arr, 0, len(arr)-1)

	tEnd := time.Now()
	times := tEnd.Sub(tStart).Seconds()
	fmt.Println(times)

	fmt.Println("排序后：", arr)
}
func randomArr() []int {
	arr:=make([]int,10000)
	for i, _ := range arr {
		arr[i] = rand.Intn(100000)
	}
	return arr
}
func QuickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	tem := arr[left]

	for i < j {
		for j > i && arr[j] >= tem {
			j--
		}
		if i < j && arr[j] < tem {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] <= tem {
			i++
		}
		if i < j && arr[i] > tem {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = tem

	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}
