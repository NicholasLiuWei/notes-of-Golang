package main

import (
	"fmt"
	"time"
	"math/rand"
)

//	优化前
func selectSort01(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

//	优化后
func selectSort02(arr []int) []int {
	i:=0
	temIntex:=i
	for ; i < len(arr)-1; i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[temIntex] > arr[j] {
				temIntex = j
			}
		}
		arr[i],arr[temIntex] = arr[temIntex],arr[i]
	}

	return arr
}

func randomArr() []int {
	arr:=make([]int,1000)
	for i, _ := range arr {
		arr[i] = rand.Intn(10000)
	}
	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())


	//arr := []int{4, 5, 7, 2, 3, 57, 8, 1, 34, 6, 756, 1, 3, 358, 4, 67, 58, 3, 1, 7, 3, 34, 3, 6, 6}
	arr := randomArr()
	fmt.Println(arr)

	tStart := time.Now()

	selectSort01(arr)
	//selectSort02(arr)

	tEnd := time.Now()
	times := tEnd.Sub(tStart).Seconds()
	fmt.Println(times)

	fmt.Println(arr)
}
