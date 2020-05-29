package main

import "fmt"

func main() {
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	students := [...]string{"Mehran", "Arman", "Iman"}
	fmt.Printf("Students: %v\n", students)

	var books [3]string

	books[0] = "Book1"
	books[1] = "Book2"
	books[2] = "Book3"

	fmt.Printf("Books: %v\n", books)
	fmt.Printf("Number of students!: %v\n", len(books))

	// defining Array matrix
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 2, 3}
	identityMatrix[1] = [3]int{4, 5, 6}
	identityMatrix[2] = [3]int{7, 8, 9}
	fmt.Printf("identityMatrix: %v\n", identityMatrix)

	// When you copy and array it basically creates a new array
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 123

	fmt.Printf("Array A: %v Array B: %v\n", a, b)
}
