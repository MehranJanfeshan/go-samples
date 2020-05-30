package main

import "fmt"

func main() {
	statePopulation := map[string]int{
		"Bavaria": 123,
		"Bremen":  122424,
		"Hamburg": 122424,
		"Hesse":   122424,
		"Saxony":  122424,
	}

	m := map[[3]int]string{}
	fmt.Println(statePopulation, m)

	// if you do not have values for your map when you create them then you can use make command to create them
	courseMap := make(map[string]int)

	courseMap["course1"] = 123
	fmt.Println(courseMap)

	courseMap = map[string]int{
		"course1": 123,
	}

	fmt.Println(courseMap)

	// adding to the map
	courseMap["course2"] = 23212
	fmt.Println(courseMap["Munich"])

	// Delete from the map
	delete(courseMap, "course2")
	fmt.Println(courseMap)

	// check the existence of ke
	_, ok := courseMap["course2"]
	fmt.Println(ok)

	// get the length
	fmt.Println(len(courseMap))

	// maps are passed by reference, so if you change a reference to a map, basically you are changing the actual map

}
