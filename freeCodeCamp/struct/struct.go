package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number    int
	actorName string
	companion []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "John Pertwee",
		companion: []string{
			"Mehran",
			"Arman",
			"Iman",
		},
	}

	fmt.Printf("This is my struct %v", aDoctor)
	defineAnonymousStruct()
	relatedStructs()
	structTag()
}

func defineAnonymousStruct() {
	aDoctor := struct {
		name   string
		number int
	}{name: "Mehran", number: 123}

	fmt.Println(aDoctor)
}

func relatedStructs() {
	type Animal struct {
		Name   string
		Origin string
	}

	type Bird struct {
		Animal
		SpeedKPH float32
		Canfly   bool
	}

	myBird := Bird{Animal: Animal{Name: "Chicken", Origin: "Earth"},
		SpeedKPH: 123,
		Canfly:   false,
	}

	fmt.Println(myBird)
}

func structTag() {
	// Tags are meaningless for GO itself but you can access them and make decision based on them!
	type Animal struct {
		Name   string `required_max:"100"`
		Origin string
	}

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
