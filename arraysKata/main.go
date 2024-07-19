package arraysKata

import "fmt"

/*
Given an integer array nums of length n, you want to create an array ans of
length 2n where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).

Specifically, ans is the concatenation of two nums arrays.

Return the array ans.
*/

func GetConcatenation(nums []int) []int {
	var newSlice = []int

	for i := 0; i <= 2; i++ {
		newSlice = append(newSlice, nums...)
	}

	return newSlice
}

func PracticeArray() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
}
