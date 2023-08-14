# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if head is None or head.next is None:
            return False

        turtle = head
        hare = head

        while hare and hare.next:
            hare = hare.next.next
            if turtle == hare:
                return True
            turtle = turtle.next

        return False

