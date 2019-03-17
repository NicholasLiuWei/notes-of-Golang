package main

import "fmt"

func main01() {

	for i := 3; i < 80000; i++ {
		var b = true
		for j := 2; j < i; j++ {
			//fmt.Printf("i:%d,j:%d,余数:%d\n",i,j,i%j)
			if i%j == 0 {
				b = false
			}

		}
		if b == true {
			fmt.Printf("%d是素数\n", i)
		}
	}
}

func putNum(numChan chan int) {
	for i := 1; i <= 8000000; i++ {
		numChan <- i
	}
	close(numChan)
}
func putPrime(primeChan chan int, numChan chan int, boolChan chan bool) {

	var flag bool
	for {
		flag = true
		num, ok := <-numChan
		if !ok {
			break
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}

	boolChan <- true
}
func main() {
	numChan := make(chan int, 80000)
	primeChan := make(chan int, 10000)
	boolChan := make(chan bool, 4)

	go putNum(numChan)

	for i := 0; i < 4; i++ {
		go putPrime(primeChan, numChan, boolChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-boolChan
		}
		close(primeChan)
	}()

	//for i := 0; i < 4; i++ {
	//	<-boolChan
	//}
	//close(primeChan)

	for {
		prime,ok := <-primeChan
		if !ok{
			break
		}
		fmt.Printf("%d是素数\n",prime)
	}
}
