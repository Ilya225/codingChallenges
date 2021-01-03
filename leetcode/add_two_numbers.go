package main

import "fmt"

// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nxt1 = l1
	var nxt2 = l2
	var remainder = 0

	root := &ListNode{}
	var nxRes = root
	var prev = root

	for nxt1 != nil && nxt2 != nil {

		var nxtVal = nxt1.Val + nxt2.Val + remainder
		if nxtVal >= 10 {
			remainder = (nxtVal) / 10
			nxRes.Val = nxtVal % 10
		} else {
			remainder = 0
			nxRes.Val = nxtVal
		}
		nxRes.Next = &ListNode{}
		prev = nxRes
		nxRes = nxRes.Next

		nxt1 = nxt1.Next
		nxt2 = nxt2.Next
	}

	prev.Next = nil

	var restList *ListNode
	if nxt2 != nil {
		restList =  nxt2
	} else if nxt1 != nil {
		restList = nxt1
	}

	if restList != nil {
		prev.Next = nxRes
		for restList != nil {
			var nxtVal = restList.Val + remainder
			if nxtVal >= 10 {
				remainder = (nxtVal) / 10
				nxRes.Val = nxtVal % 10
			} else {
				nxRes.Val = nxtVal
				remainder = 0
			}
			restList = restList.Next
			nxRes.Next = &ListNode{}
			prev = nxRes
			nxRes = nxRes.Next
		}
	}

	if remainder != 0 {
		if prev.Next == nil {
			prev.Next = &ListNode{}
		}
		prev.Next.Val += remainder
	} else {
		prev.Next = nil
	}

	return root
}


func addTwoNumbersApp() {
	var l1 = createListNode([]int{2,4,3})
	var l2 = createListNode([]int{5,6,4})

	fmt.Printf("%v", l1)
	fmt.Printf("%v", l2)

	var r = addTwoNumbers(l1, l2)

	fmt.Printf("%v", r)
}

func createListNode(nums []int) *ListNode {
	root := &ListNode{}
	var nextNode = root
	var prevNode = root
	for i:= 0; i < len(nums); i++ {
		nextNode.Val = nums[i]
		prevNode = nextNode
		nextNode.Next = &ListNode{}
		nextNode = nextNode.Next
	}

	prevNode.Next = nil

	return root
}