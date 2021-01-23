package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func generator3(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:

			case <-done:
				return
			}

		}

	}()
	return out
}

func square3(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:

			case <-done:
				return
			}
		}
	}()
	return out
}

func merge3(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
				case out <- n:

				case <-done:
					return
			}

		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {

	done := make(chan struct{})
	in := generator3(done,2, 3)

	c1 := square3(done, in)
	c2 := square3(done, in)

	out := merge3(done, c1, c2)

	fmt.Println(<-out)
	close(done)

	time.Sleep(10 * time.Millisecond)
	fmt.Println(runtime.NumGoroutine())
}