package main

import "fmt"

func WriteData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Printf("写入数据：%d\n", i)
	}
	close(intChan)
}
func ReadData(intChan chan int, boolChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("读取数据%d\n", v)
	}
	boolChan <- true
	close(boolChan)
}

func main() {
	intChan := make(chan int, 50)
	boolChan := make(chan bool)

	go WriteData(intChan)
	go ReadData(intChan, boolChan)

	for {
		_, ok := <-boolChan
		if ok {
			break
		}
	}
}
