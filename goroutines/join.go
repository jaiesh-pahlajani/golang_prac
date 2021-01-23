package main

import (
	"fmt"
	"sync"
)

func main() {
	// Wait group to wait for go routine to finish
	var wg sync.WaitGroup
	wg.Add(1)

	var data int

	go func() {
		data++
		wg.Done()
	}()

	wg.Wait()

	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}