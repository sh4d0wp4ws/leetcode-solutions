package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int { // This function received 2 input arguments are "nums" and "target"
	for i := 0; i <= len(nums)-1; i++ {
		for y := i + 1; y <= len(nums)-1; y++ {
			if nums[i]+nums[y] == target {
				return []int{i, y}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target)) // Print the result of the twoSum function
}
