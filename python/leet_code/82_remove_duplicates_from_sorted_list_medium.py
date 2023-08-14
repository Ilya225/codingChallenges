# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:

        first = head
        pivot = None
        root = None

        while first and first.next:
            next_node = first.next

            if first.val != next_node.val:
                if not root:
                    root = first

                if pivot:
                    pivot.next = first
                pivot = first
                pivot.next = None

            while next_node and first.val == next_node.val:
                first = next_node
                next_node = next_node.next

            first = next_node

        if first:
            if not root:
                root = first
            else:
                pivot.next = first

        return root
