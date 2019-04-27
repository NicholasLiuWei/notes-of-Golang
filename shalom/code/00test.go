package main

import (
	"fmt"
)

func main() {
	a := &struct{}{}
	b := &struct{}{}
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", b)

	fmt.Println(a == b)
}
