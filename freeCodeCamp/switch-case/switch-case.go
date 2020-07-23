package main

import "fmt"

func main() {
	switch 2 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Not one or two")
	}

	orSwitchCase()
	fallthroughSwitchCase()
	typeSwitch()
	exitWithBreakSwitch()
}

func orSwitchCase() {
	switch 2 {
	case 1, 3, 5:
		fmt.Println("One, Three, Five")
	case 2, 4, 6:
		fmt.Println("Two, Four, Six")
	default:
		fmt.Println("Another number!")
	}
}

func fallthroughSwitchCase() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("One, Three, Five")
		// by default in GO only one case executes if the condition satisfied but if more than one case is needed to be executed then you can use
		// fallthrough keyword
		fallthrough
	case i <= 20:
		fmt.Println("Two, Four, Six")
	default:
		fmt.Println("Another number!")
	}
}

func typeSwitch() {
	var i interface{} = [3]int{}
	switch i.(type) {
	case int:
		fmt.Println("i is int")
	case float64:
		fmt.Println("i is float64")
	case string:
		fmt.Println("i is string")
	case [3]int:
		fmt.Println("i is [3]int")
	default:
		fmt.Println("unknown type!")
	}
}

func exitWithBreakSwitch() {
	switch 1 {
	case 1:
		fmt.Println("before break!")
		break
		fmt.Println("after break!")
	case 2:
		fmt.Println("Tow")
	case 3:
	default:
		fmt.Println("unknown value!")
	}
}
