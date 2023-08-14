# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def pairSum(self, head: Optional[ListNode]) -> int:

        list_len = 0
        cur = head
        while cur:
            list_len += 1
            cur = cur.next

        midddle = list_len // 2


        max_sum = 0
        index = 0
        prev = None
        cur = head
        start = head
        while index < list_len:
            if index < midddle:
                cur = cur.next
            else:
                next_node = cur.next
                cur.next = prev
                prev = cur
                cur = next_node
            index += 1


        for i in range(0, midddle):
            max_sum = max(max_sum, prev.val + start.val)
            start = start.next
            prev = prev.next

        return max_sum

