package main

import "fmt"

func main() {
	arr:=[]int{4,5,7,2,3,57,8,1,34,6,756,1,3,358,4,67,58,3,1,7,3,34,3,6,6}

	//	第一层循环，每循环一次排好一个值
	for i := 0; i < len(arr); i++ {
		//	第二层循环，将无序的部分相邻两两作比较，冒出最大值在最后边
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}

	fmt.Println(arr)
}
