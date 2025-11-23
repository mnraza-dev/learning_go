package main

import "fmt"

func main() {
	// age := 8
	// if age >= 18 {
	// 	println("You are an adult.")
	// } else {
	// 	println("You are a minor.")
	// }
	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// } else if age >= 13 {
	// 	fmt.Println("You are a teenager.")
	// } else {
	// 	fmt.Println("You are a child.")
	// }

	// var role string = "admin"
	// var hasPermission bool = false
	// if role == "admin" || hasPermission {
	// 	fmt.Println("Access granted.")
	// }
	// if role == "admin" && hasPermission {
	// 	fmt.Println("Access granted.")
	// }

	// // short statement in if
	if age := 20; age >= 18 {
		fmt.Println("You are an adult.", age)
	} else if age >= 13 {
		fmt.Println("You are a teenager.", age)
	} else {
		fmt.Println("You are a child.", age)
	}

	// Go doesn't have ternary operator
	

}
