package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		fmt.Println("This is a old number")
		i := <-ch
		fmt.Println(i)
		ch <- 25
		wg.Done()
	}()

	go func() {
		ch <- 42
		fmt.Println("This is a new number", <-ch)
		wg.Done()
	}()

	wg.Wait()
}
