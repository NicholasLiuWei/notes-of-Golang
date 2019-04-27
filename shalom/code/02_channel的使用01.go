package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	m    = make(map[int]int, 10)
	lock = sync.Mutex{}
)

func Cal(i int) {

	lock.Lock()
	var res = 1
	for j := 1; j <= i; j++ {
		res *= j
	}
	m[i] = res
	lock.Unlock()
}

func main() {

	for i := 1; i <= 20; i++ {
		go Cal(i)
	}
	time.Sleep(10 * time.Second)

	lock.Lock()
	for i, v := range m {
		fmt.Printf("m[%d!] = %d\n", i, v)
	}
	lock.Unlock()
}
