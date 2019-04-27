package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binarySearch(arr []int, aim int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		fmt.Println(arr[mid])

		if aim < arr[mid] {
			right = mid - 1
		} else if aim > arr[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
func randomArr() []int {
	arr := make([]int, 1000)
	for i, _ := range arr {
		arr[i] = rand.Intn(10000)
	}
	return arr
}
func main() {
	rand.Seed(time.Now().UnixNano())
	arr := randomArr()
	arr = append(arr, 8620)

	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	result := binarySearch(arr, 8620)
	if result != -1 {
		fmt.Println("查找结果: ", result)
	} else {
		fmt.Println("数据中没有目标值")
	}
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