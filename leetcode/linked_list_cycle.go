package main

import "fmt"

// https://leetcode.com/problems/linked-list-cycle-ii/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	turtle := head
	hare := head

	if hare.Next != nil && hare.Next.Next != nil {
		turtle = turtle.Next
		hare = hare.Next.Next
	} else {
		return nil
	}

	for turtle!= hare {
		turtle = turtle.Next
		if hare.Next != nil && hare.Next.Next != nil {
			hare = hare.Next.Next
		} else {
			return nil
		}
	}


	turtle = head
	for turtle != hare {
		turtle = turtle.Next
		hare = hare.Next
	}

	return hare
}

func detectCycleApp() {
	var list []*ListNode
	for _, n := range []int{3,2,0,-4} {
		list = append(list, &ListNode{Val: n})
	}

	for i:= 0; i < len(list)-1; i++ {
		list[i].Next = list[i+1]
	}

	list[len(list)-1].Next = list[1]

	cycle := detectCycle(list[0])
	fmt.Printf("%v\n", cycle)
}
