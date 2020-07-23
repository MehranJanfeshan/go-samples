package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.google.com/robots.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)

	deferSample()
}

func deferSample() {
	// defer always get the values at the time that defer declared

	a := "start"
	defer fmt.Println(a)
	a = "end"

	// it prints out start
}
