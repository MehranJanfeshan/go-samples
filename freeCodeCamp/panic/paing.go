package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start!")
	paniker()
	fmt.Println("second!")
}


// defer always runs after panic and before function ends
// Panic stops every code after itself but only within the function that is called
// recover can be used to recover from panic and let the execution of program continues

func paniker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Print("Error:", err)
		}
	}()

	panic("Something bad happened!")
	fmt.Println("done Paniking")
}
