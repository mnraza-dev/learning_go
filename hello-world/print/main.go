package main

import "fmt"

func main() {
	var age int = 30
	var name string = "MN Raza"
	height := 5.57473

	fmt.Println("Age :", age, "Name: ", name, "Height :", height)
	fmt.Printf("Age: %d, Name: %s, Height: %.2f\n", age, name, height)
	
	fmt.Printf("Type of age: %T, Type of name: %T, Type of height: %T\n", age, name, height)
	
}
