package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	Name string
	Age int
	Gender string
} 

func reflect01(i interface{})  {
	//	转成rType类型
	rType:=reflect.TypeOf(i)
	fmt.Println(rType)
	//fmt.Printf("%T",rType)
	
	//	获得rValue类型
	rVal:=reflect.ValueOf(i)
	fmt.Println(rVal)
	fmt.Println(rVal.Int())

	//	重新转回原数据
	num2:=rVal.Interface().(int)
	fmt.Println(num2)
}

func reflect02(i interface{})  {
	//	转成rType类型
	rType:=reflect.TypeOf(i)
	fmt.Println(rType)
	//fmt.Printf("%T",rType)

	//	获得rValue类型
	rVal:=reflect.ValueOf(i)
	fmt.Println(rVal)

	//	重新转回原数据
	iv:=rVal.Interface()
	//	iv运行前为接口类型，不能够取出结构体的属性，运行时为结构体类型
	fmt.Printf("%v  type of %T\n",iv,iv)
}

func main() {
	var num = 100
	reflect01(num)

	student:=Student{
		"shalom",
		19,
		"male",
	}
	reflect02(student)
}

