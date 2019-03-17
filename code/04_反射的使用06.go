package main

import (
	"reflect"
	"fmt"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (this Cal) GetSub(name string) {
	res := this.Num1 - this.Num2
	fmt.Printf("%s完成了减法运行，%d - %d = %d", name, this.Num1, this.Num2, res)
}

func reflect01(pem interface{}) {
	rVal := reflect.ValueOf(pem)
	num := rVal.Elem().NumField()
	for i := 0; i < num; i++ {
		vi := rVal.Elem().Field(i)
		fmt.Printf("第%d个字段的值：%d\n", i, vi)
	}

	var para []reflect.Value
	para = append(para,reflect.ValueOf("tom"))
	rVal.Elem().Method(0).Call(para)
}

func main() {
	var cal = Cal{8, 3}

	reflect01(&cal)

}
