package main

import "fmt"

func main() {
	ifWithInitializer()
}

func ifWithInitializer() {
	courses := map[string]int{
		"Course1": 2342,
		"Course2": 2342,
	}

	// this part called initializer: course, ok := courses["Course1"];
	if course, ok := courses["Course1"]; ok {
		fmt.Println(course)
	}
}
