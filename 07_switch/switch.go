package main

import (
	"fmt"
)

func main() {
	// simple switch case
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("i is 1")
	// case 2:
	// 	fmt.Println("i is 2")
	// case 3:
	// 	fmt.Println("i is 3")
	// default:
	// 	fmt.Println("i is something else")
	// }

	// multiple condition  in switch
	// switch time.Now().Weekday() {
	// case time.Monday:
	// 	fmt.Println("Office ke liye Taiyyar ho jao!")
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend!")
	// default:
	// 	fmt.Println("It's a weekday.")
	// }

	// type switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a boolean")
		default:
			fmt.Println("I am of a different type", t)
		}
	}

	whoAmI(42)
	whoAmI("Hello, Go!")
	whoAmI(true)
	whoAmI(3.14)

}
