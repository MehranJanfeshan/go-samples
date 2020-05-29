package main

import "fmt"

// slices are referred by reference!
func main() {
	a := []int{1, 2, 3}
	b := a
	b[0] = 123

	fmt.Printf("A %v\n", a)
	fmt.Printf("B %v\n", b)
	fmt.Printf("Length %v\n", len(a))
	fmt.Printf("Length %v\n", cap(a))

	// slicing
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]   // slice of all elements
	e := c[3:]  // slice from 4th element to end
	f := c[:6]  // slice first 6 elements
	g := c[3:6] // slice the 4th, 5th and 6th elements

	fmt.Printf("C %v\n", c)
	fmt.Printf("D %v\n", d)
	fmt.Printf("E %v\n", e)
	fmt.Printf("F %v\n", f)
	fmt.Printf("G %v\n", g)

	// using make
	h := make([]int, 3, 100)
	fmt.Println(h)
	fmt.Printf("Length %v\n", len(h))
	fmt.Printf("Capacity %v\n", cap(h))

	h = append(h, 1)
	fmt.Println(h)
	fmt.Printf("Length %v\n", len(h))
	fmt.Printf("Capacity %v\n", cap(h))

	// Declaring array without init capacity va values
	var j []int
	fmt.Println(j)
	fmt.Printf("Length %v\n", len(j))
	fmt.Printf("Capacity %v\n", cap(j))

	j = append(j, 1)
	fmt.Println(j)
	fmt.Printf("Length %v\n", len(j))
	fmt.Printf("Capacity %v\n", cap(j))

	// appending to array using spread function
	j = append(j, []int{1, 2, 3, 4}...)
	fmt.Println(j)

	// removing items from slice

	k := []int{1, 2, 3, 4, 5}
	l := append(k[:2], k[3:]...)
	fmt.Println(l)

}
