package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Hello, World from main package!")
	myutil.Hello("Hello, World!")

	var name string = "John Doe"
	var version string = "1.0.0"
	var money int = 61200
	var isSelected bool = true
	const PI = 3.14

	fmt.Println(PI)
	fmt.Println(isSelected)
	fmt.Println(money)
	fmt.Println("Hello, " + name + "! Version: " + version)

	// without using var and const
	person := "Jane Doe"
	fmt.Println("Hello, " + person + "!")

	var Public = "Data is imp"           // can be used outside the package
	var private = "Data is very private" // used only inside the package
	fmt.Println(Public)
	fmt.Println(private)

	
}
