package main

import "fmt"

// https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := []int{}

	var median float64
	var i = 0
	var j = 0

	for i < len(nums1) && j< len(nums2) {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	if i<len(nums1) {
		res = append(res, nums1[i:]...)
	} else if (j < len(nums2)) {
		res = append(res, nums2[j:]...)
	}

	if (len(res) % 2 == 0) {
		median = float64(res[len(res)/2] + res[len(res)/2-1])/2.0
	} else {
		median = float64(res[len(res)/2])
	}

	return median
}

func findMedianSortedArraysApp() {
	var median = findMedianSortedArrays([]int{2}, []int{})
	fmt.Println(median)
}
