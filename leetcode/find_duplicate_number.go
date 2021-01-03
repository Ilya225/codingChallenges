package main

import "fmt"

// https://leetcode.com/problems/find-the-duplicate-number/

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		idx := nums[i]
		if nums[i] < 0 {
			idx = -nums[i]
		}
		if nums[idx-1] < 0 {
			return idx
		}
		nums[idx-1] = -nums[idx-1]
	}

	return 0
}

func findDuplicateApp() {
	fmt.Printf("%v", findDuplicate([]int{1}))
}
