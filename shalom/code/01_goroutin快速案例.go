package main

import (
	"fmt"
	"strconv"
	"time"
	"runtime"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	//go test()
	//for i := 0; i < 10; i++ {
	//	fmt.Println("hello golang" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}
	num:=runtime.NumCPU()
	fmt.Println(num)

	runtime.GOMAXPROCS(num)
}
