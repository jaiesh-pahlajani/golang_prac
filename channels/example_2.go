package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
		}
		close(ch)
	}()

	for elem := range ch {
		fmt.Println(elem)
	}

}
