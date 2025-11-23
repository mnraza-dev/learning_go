package main

import "fmt"

func main() {

	// slices -> dynamic arrays
	// - most used construct in go + useful methods

	// uninitialized slice is nill
	// var nums []int             - nil
	// fmt.Println("Slice:", nums)
	// fmt.Println(nums == nil)
	// fmt.Println("Length:", len(nums))

	var nums = make([]int, 5, 6) // initialize slice with size 5
	fmt.Println("Slice:", nums)
	fmt.Println("Length:", len(nums))
	fmt.Println("Capacity", cap(nums))
	nums = append(nums, 10)
	nums = append(nums, 23)
	nums = append(nums, 67)
	nums = append(nums, 234)
	nums = append(nums, 674)
	nums = append(nums, 637)
	nums = append(nums, 6743)
	nums = append(nums, 6738)
	nums = append(nums, 6373)
	fmt.Println("Slice after append:", nums)
	fmt.Println("Length after append:", len(nums))
	fmt.Println("Capacity after append:", cap(nums))

}
