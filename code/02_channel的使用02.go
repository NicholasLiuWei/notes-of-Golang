package main

import "fmt"

func main() {

	var ch01 chan int
	//var ch02 chan map[string]int

	ch01 = make(chan int, 3)
	ch01 <- 10
	num := <-ch01
	fmt.Println(num)

	var i interface{}
	i = 19
	n:=i.(bool)
	fmt.Println(n)
}
