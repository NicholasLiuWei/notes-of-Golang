package main

import "fmt"

func putNum(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func cal(numChan chan int, resChan chan map[int]int, boolChan chan bool) {
	mres := make(map[int]int)
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		var res int

		for i := 0; i <= num; i++ {
			res += i
		}
		mres[num] = res
	}
	resChan <- mres
	boolChan <- true
}

func main() {
	numChan := make(chan int, 2000)
	resChan := make(chan map[int]int, 10)
	boolChan := make(chan bool, 8)

	go putNum(numChan)
	for i := 0; i < 8; i++ {
		go cal(numChan, resChan, boolChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-boolChan
		}
		close(resChan)
	}()

	for {
		res := <-resChan
		for i, v := range res {
			fmt.Printf("res[%d]=%d\n", i, v)
		}
	}
}
