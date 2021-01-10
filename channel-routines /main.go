package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://www.google.com/",
		"https://de.yahoo.com/",
		"https://www.test.com/",
		"https://www.barnesandnoble.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}

	// infinity loop
	//for {
	//	go checkLink(<-c, c)
	//}

	// infinity loop but with better syntax
	//for l := range c {
	//	// This is a blocking and you should not do it as it is against the go routine!
	//	time.Sleep(2 * time.Second)
	//	go checkLink(l, c)
	//}

	// This is a infinity loop
	// The difference between this loop and other is that this loop has pause time within another function
	// it means the whole loop wont be blocked
	for l := range c {
		go func(l string) {
			time.Sleep(2 * time.Second)
			checkLink(l, c)
		}(l)

	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	c <- link
	fmt.Println(link, "is up!")
}
