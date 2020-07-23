package main

import "fmt"

func main() {
	// using increment values!
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for without condition - it can be infinity
	fmt.Println("for without condition - it can be infinity")
	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}

	// using continue to iterate earlier

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// break from the inner loop
	fmt.Println("break from the inner loop")
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// loop over array
	fmt.Println("loop over slice")
	s := [3]int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	// loop over map
	fmt.Println("loop over map")
	statePopulation := map[string]int{
		"California": 2134234,
		"Texas":      1234123,
		"Florida":    123123,
	}

	for k, v := range statePopulation {
		fmt.Println(k, v)
	}

	// loop over string
	str := "Hello GO"
	for k, v := range str {
		fmt.Println(k, v)
	}
}
