package main

import "fmt"

var counter = 0

func main() {
	//var result = fib(50)
	//fmt.Print(result)
	recursive1(5)
	fmt.Println(counter)
}

func fib(num int64) int64 {
	if num <= 2 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}

func recursive1(n int64) {
	counter += 1
	if n <= 1 {
		return
	}
	recursive1(n - 1)
	recursive1(n - 1)
}
