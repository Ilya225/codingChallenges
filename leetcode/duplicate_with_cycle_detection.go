package main

// https://leetcode.com/problems/find-the-duplicate-number/

func findDuplicateCycle(nums []int) int {
	turtle := nums[0]
	hare := nums[0]

	turtle = nums[turtle]
	hare = nums[nums[hare]]

	for turtle != hare {
		turtle = nums[turtle]
		hare = nums[nums[hare]]
	}

	turtle = nums[0]
	for turtle != hare {
		turtle = nums[turtle]
		hare = nums[hare]
	}

	return hare

}


