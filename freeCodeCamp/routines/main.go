package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}
// The reason that we use lock is to make sure the race condition
// on goroutines is not happening, With the lock we can make sure
// that we only read and write one at the time
// This is really a bad practice because we are basically stopping concurrency
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // read lock
		go sayHello()
		m.Lock() // write lock
		go increment()

	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
