package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"hhh"`
	Age  int    `json:"doubi"`
}

func main() {
	shalom := Person{
		"shalom",
		19,
	}
	fmt.Println(shalom)
	data, err := json.Marshal(shalom)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))
}
