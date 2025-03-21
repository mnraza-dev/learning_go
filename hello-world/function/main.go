package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}
func sum(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("Learning functions in Golang...")
	hello()
	fmt.Println("Sum of two number is : ", sum(19, 2))
}
