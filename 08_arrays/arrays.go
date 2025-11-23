package main

import "fmt"

func main() {

	// var nums [4]int
	// nums[0] = 10
	// fmt.Println("Array:", nums)
	// fmt.Println("Array:", len(nums))
	// fmt.Println("First element:", nums[0])

	// var vals [4]bool
	// vals[2] = true
	// fmt.Println(vals)

	// var fruits [3]string
	// fruits[0] = "Apple"
	// fmt.Println("Fruits:", fruits)



	// to declare it in  single line
	// nums := [4]int{10, 20, 30, 40}
	// fmt.Println("Numbers:", nums)
	// fmt.Println("Length:", len(nums))

	// 2d array
	nums := [2][2]int{
		{1, 2},
		{4, 6},
	}
	fmt.Println("2D Array:", nums)
}
