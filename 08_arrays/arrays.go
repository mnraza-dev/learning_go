package main

import "fmt"

func main() {

	var nums [4]int
	nums[0] = 10
	fmt.Println("Array:", nums)
	fmt.Println("Array:", len(nums))
	fmt.Println("First element:", nums[0])
}
