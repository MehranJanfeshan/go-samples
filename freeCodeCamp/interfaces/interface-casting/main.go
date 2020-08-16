package main

import "fmt"

func main() {
	// If your instruct method has pointer receiver then you need to & to assign instruct to a interface.
	// By using & you can have pointer and value receiver methods
	var wc WriterCloser = &myWriterCloser{}
	fmt.Println(wc)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct{}

func (mwc *myWriterCloser) Write([]byte) (int, error) {
	return 0, nil
}

func (mwc *myWriterCloser) Close() error {
	return nil
}
