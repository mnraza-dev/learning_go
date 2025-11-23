package main

import "fmt"

func main() {

	//for loop -> only construct in go for iteration

	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// infinite loop
	// for{
	// 	fmt.Println("Hello, Go developers!")
	// }

	// traditional for loop
	// for i := 1; i <= 3; i++ {
	// 	fmt.Println(i)
	// }

	// continue statement
	// for i := 1; i <= 3; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	// break statement
	// for i := 1; i <= 3; i++ {
	// 	if i == 2 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// range keyword
	for i := range 11 {
		fmt.Println(i)
	}
}
