package main

import (
	"fmt"
)

func main() {
	// The iota keyword represents successive integer constants 0, 1, 2,…
	const (
		C1 = iota
		C2
		C3
	)

	fmt.Println(C1, C2, C3) // "0, 1 2"

	const (
		_ = iota
		C4
		C5
		C6
	)

	fmt.Println(C4, C5, C6) // "1 2 3"

	const (
		_  = iota
		KB = 1 << (10 * iota) // 1 * 2^10
		MB                    // 1 * 2^100
	)

	fmt.Println(KB, MB) // "1 2 3"

	const (
		isAdmin = 1 << iota
		catSpecialist
		dogSpecialist
		snakeSpecialist
	)

}
