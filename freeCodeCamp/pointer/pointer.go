package main

import "fmt"

func main() {
	simplePointer()
	simplePointer2()
	simplePointer3()
	simplePointer4()
	slicePointer()
	mapPointer()
}

// pointer to variables
func simplePointer() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(&a, b)
	fmt.Println(a, *b)
}

// pointer to the array elements
func simplePointer2() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v, %p %p\n", a, b, c)
}

// pointer to the struct
type myStruct struct {
	foo int
}

func simplePointer3() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)
}

// dereference the pointer for struct!
func simplePointer4() {
	var ms *myStruct
	ms = new(myStruct)
	(*ms).foo = 123
	fmt.Println(ms.foo)
}

// slice is a pointer to the first element of array so bay changing a slice other references to the slice also get updated!
func slicePointer() {
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[0] = 123
	fmt.Println(a, b)
}

// map has the same behavior as slice so need to be careful who has access to the map as changing the map could affect other parts of the code.
func mapPointer() {
	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "QUYX"
	fmt.Println(a, b)

}
