package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"年龄"`
	Gender string
}



func reflect01(i interface{}) {
	val := reflect.ValueOf(i)
	typ := reflect.TypeOf(i)
	num := val.NumField()
	fmt.Println(num)

	for i := 0; i < num; i++ {
		fmt.Printf("第%d个字段的值：%v\n",i,val.Field(i))
		tagVal:=typ.Field(i).Tag.Get("json")
		if tagVal!=""{
			fmt.Printf("第%d个字段的标签：%v\n",i,tagVal)
		}
	}
}

func main() {
	stu := Student{"shalom", 19, "male"}

	reflect01(stu)
}
