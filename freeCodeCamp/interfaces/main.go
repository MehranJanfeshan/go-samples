package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Here we use interface to access the write and close function
	var wc WriterCloser = NewBufferWriterCloser()
	wc.Write([]byte("Hello YouTube listener, this is a test"))
	wc.Close()

	// In case if you want to cast the wc to a struct in order to access
	// internal variables or any functions that are not defined in interface you can
	// do the cast by below command!
	bwc := wc.(*BufferWriterCloser)
	fmt.Println(bwc)

	// To have a safe casting you can use the below command:
	SafeBwc, ok := wc.(*BufferWriterCloser)
	if ok {
		fmt.Println(SafeBwc)
	}
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

type BufferWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)

	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferWriterCloser() *BufferWriterCloser {
	return &BufferWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
