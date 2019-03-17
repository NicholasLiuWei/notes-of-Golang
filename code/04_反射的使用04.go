package main

import (
	"reflect"
	"fmt"
)



func reflect01(i interface{})  {
	rVal:=reflect.ValueOf(i)

	ty:=rVal.Type()
	ki:=rVal.Kind()
	fmt.Println(ty,ki)

	v02:=rVal.Interface().(float64)
	fmt.Println(v02)
}

func main() {
	var v float64 = 1.2
	reflect01(v)
}

