package main

import "fmt"

// https://leetcode.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < 1 || nums[i] == 0 {
			nums[i] = len(nums)+1
		}
	}

	for i := 0; i < len(nums); i++ {
		idx := nums[i]
		if nums[i] < 0 {
			idx = -nums[i]
		}

		if idx <= len(nums) && nums[idx-1] > 0 {
			nums[idx-1] = -nums[idx-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i+1;
		}
	}
	return len(nums)+1;
}

func firstMissingPositiveApp() {
	fmt.Printf("%v", firstMissingPositive([]int{1,1}))
}
