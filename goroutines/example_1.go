package main

import (
	"fmt"
	"time"
)

func funCall(value string)  {
	for i := 0; i < 3; i++ {
		fmt.Println(value)
		time.Sleep(1 * time.Millisecond)
	}
}

func main()  {
	// Direct call
	funCall("Direct call")

	// Go routine fun call
	go funCall("go routine 1")

	// Go routine with anon func
	go func() {
		funCall("go routine 2")
	}()

	// Go routine with function value call
	fv := funCall
	go fv("go routine 3")

	// wait for routines to end
	time.Sleep(100 * time.Millisecond)
}
