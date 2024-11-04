package main

import "fmt"

func average(nums ...int) {
	if len(nums) == 0 {
		fmt.Println("No numbers provided")
		return

	}

	total := 0
	for _, num := range nums {
		total += num
	}

	avg := float64(total) / float64(len(nums))
	// fmt.Println("Average:", avg)
	fmt.Printf("Average: %.2f\n", avg) // %.2f formats the float to 2 decimal places
}

func main2() {

	average(1, 2)
	average(1, 2, 3)
	average(1, 2, 3, 4)
	average()

	nums := []int{1, 2, 3, 4}
	average(nums...)
}
