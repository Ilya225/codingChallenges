from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:

        first = head

        while first and first.next:
            second = first.next
            if first.val == second.val:
                first.next = second.next
            else:
                first = second

        return head
