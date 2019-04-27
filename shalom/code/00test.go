package main

import (
	"fmt"
	"reflect"
)

type Cal01 struct {
	Num1 int
	Num2 int
}

func (this *Cal01) GetSub01(name string) {
	res := this.Num1 - this.Num2
	fmt.Printf("%s完成了减法运行，%d - %d = %d", name, this.Num1, this.Num2, res)
}
func reflect03(pem interface{}) {
	rVal := reflect.ValueOf(pem)
	var para []reflect.Value
	para = append(para,reflect.ValueOf("tom"))
	rVal.Method(0).Call(para)
}
func main() {
	var cal = Cal01{8, 3}
	reflect03(&cal)
}
