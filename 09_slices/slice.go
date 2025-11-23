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

	// var nums = make([]int, 0, 6) // initialize slice with size 5
	// nums := []int{} // initialize empty slice
	// fmt.Println("Slice:", nums)
	// fmt.Println("Length:", len(nums))
	// fmt.Println("Capacity", cap(nums))
	// nums = append(nums, 10)
	// nums = append(nums, 23)
	// nums = append(nums, 67)
	// fmt.Println("Slice after append:", nums)
	// fmt.Println("Length after append:", len(nums))
	// fmt.Println("Capacity after append:", cap(nums))

	// copying slices
	// 	newNums := make([]int, len(nums), cap(nums)*2)
	// 	fmt.Println(nums, newNums)
	// 	copy(newNums, nums)
	// 	fmt.Println("New Slice:", newNums)
	// 	fmt.Println("Length of new slice:", len(newNums))
	// 	fmt.Println("Capacity of new slice:", cap(newNums))

	// slice operators
	var nums = []int{1, 2, 3, 4}
	fmt.Println("Original Slice:", nums)
	fmt.Println("Slice from index 0 to 2:", nums[0:2]) // [1 2 ]

}
