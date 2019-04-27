package main

import "fmt"

func main() {
	arr:=[]int{4,5,7,2,3,57,8,1,34,6,756,1,3,358,4,67,58,3,1,7,3,34,3,6,6}

	fmt.Println(arr)
	//	第一层遍历，开拓疆土
	for i := 0; i < len(arr); i++ {
		var j int
		//	临时变量存储第一层开拓到的新元素
		tem := arr[i]
		//	第二层遍历，将所有比tem大的元素后移一格，最终找到没有比tem大的元素，再插入
		//	注意: 当不满足for循环的中间条件时，不会进行 j-- 操作
		for j = i; j > 0 && arr[j-1] > tem; j-- {
			//	因为arr[j-1]的值比tem大，所以把它的值移到arr[j]的位置，最开始arr[j]是tem的位置
			arr[j] = arr[j-1]
		}
		//	找到tem的位置，对它进行赋值
		arr[j] = tem
	}
	fmt.Println(arr)
}
