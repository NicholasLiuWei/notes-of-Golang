package main

import (
	"reflect"
	"fmt"
)

//通过反射，修改 num int的值

func reflect01(i interface{})  {
	
	//	获得rValue类型
	rVal:=reflect.ValueOf(i)
	fmt.Println(rVal)

	//	rVal.Elem()获取rVal代表的指针指向的值
	rVal.Elem().SetInt(20)
	//fmt.Println(rVal)
}

func main() {
	var num = 100
	reflect01(&num)
	fmt.Println(num)
}

