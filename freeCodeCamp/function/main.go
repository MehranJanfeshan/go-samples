package main

import "fmt"

func main() {
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	sum(1, 2, 3, 4, 5)
	result := returnPointer()
	println(*result)
	result2 := nameReturn()
	println(result2)

	_, err := returnError(0.0, 2.0)

	if err != nil {
		fmt.Println(err)
	}
}

// pointer and function
func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Test"
	fmt.Println(*name)
}

// combines all the param into an array
func sum(values ...int) {
	for i, v := range values {
		fmt.Printf("%v, %v\n", i, v)
	}
}

// return pointer
func returnPointer() *int {
	value := 123
	return &value
}

// name return value, in this case you do not need return the value just use the return keyword.

func nameReturn() (result int) {
	result = 456
	return
}

// return error

func returnError(a float32, b float32) (float32, error) {
	if a == 0.0 {
		return 0.0, fmt.Errorf("a is 0")
	}
	return a * b, nil
}
