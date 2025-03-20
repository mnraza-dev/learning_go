package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Hello, World from main package!")
	myutil.Hello("Hello, World!")

	var name string = "John Doe"
	var version = "1.0.0"
	fmt.Println("Hello, " + name + "! Version: " + version)
}
