package main

import (
	"fmt"
	"strconv"
)

func main() {
	i:= 123
	// convert int to float
	j:= float32(i)

	// convert int to string
	s:= strconv.Itoa(i)

	// v stand for value
	// T stand for type
	fmt.Printf("%T, %T, %T", i, j, s)
	fmt.Printf("%v, %v, %v", i, j, s)


	// The iota keyword represents successive integer constants 0, 1, 2,…
	const (
		C1 = iota
		C2
		C3
	)

	fmt.Println(C1, C2, C3) // "1 2 3"
}
