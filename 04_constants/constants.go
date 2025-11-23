package main

import "fmt"

const age int = 30

// name := "Go Developer" // invalid outside function
func main() {

	// const name string = "MN Raza"
	// // name := "Go Developer"
	// fmt.Println("Hello, " + name + "!")
	fmt.Println("Age:", age)

	// grouping of constants
	const (
		city    = "Bangalore"
		country = "India"
	)

	fmt.Println("City:", city)
	fmt.Println("Country:", country)

}
