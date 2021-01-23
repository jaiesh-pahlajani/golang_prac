package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	deposit := func(amount int) {
		mutex.Lock()
		balance += amount
		mutex.Unlock()
	}

	withdrawal := func(amount int) {
		mutex.Lock()
		balance -= amount
		mutex.Unlock()
	}

	// make 100 deposits of $1
	// and 100 withdrawal of $1 concurrently.

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}

	wg.Wait()
	fmt.Println(balance)
}