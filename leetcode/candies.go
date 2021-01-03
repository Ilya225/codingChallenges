package main

import (
	"fmt"
)

// https://leetcode.com/problems/candy/

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	} else if len(ratings) == 1 {
		return 1
	}

	candies := make([]int, len(ratings))
	prev := 0
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i-1] < ratings[i] {
			prev += 1
		} else {
			prev = 1
		}
		candies[i] = prev
	}

	prev = 0
	for i := len(ratings)-1; i >= 0; i-- {
		if i < len(ratings)-1  && ratings[i+1] < ratings[i] {
			prev += 1
		} else {
			prev = 1
		}
		if candies[i] < prev {
			candies[i] = prev
		}
	}

	res := 0
	for _,e := range candies {
		res += e
	}
	return res
}

func candyWalk(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	} else if len(ratings) == 1 {
		return 1
	}
	counter := 1
	if ratings[1] < ratings[0] {
		counter += 1
	}

	res := 0
	for i := 0; i < len(ratings); i++ {
		if (i < len(ratings)-1 && i > 0 && ratings[i] == ratings[i+1] && ratings[i-1] == ratings[i]) || (i == len(ratings)-1 && ratings[i-1] == ratings[i]) {
			res += 1
			continue
		}

		if (i < len(ratings)-1 && ratings[i] > ratings[i+1]) && (i > 0 && ratings[i-1] == ratings[i]) {
			if counter == 1{
				counter += 1
			}
		}

		res += counter

		if i < len(ratings)-1 && ratings[i] < ratings[i+1] {
			counter += 1
		} else if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			if counter > 1 {
				counter -= 1
			}
			if i == len(ratings)-2 {
				counter = 1
			}
		}
	}
	return res
}

func candyApp() {
	fmt.Println(candyWalk([]int{1,3,4,5,2}))
}
