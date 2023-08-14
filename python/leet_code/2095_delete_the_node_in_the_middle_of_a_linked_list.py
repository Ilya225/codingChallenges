from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        list_len = 0

        cur = head
        while cur:
            list_len += 1
            cur = cur.next

        middle = list_len // 2
        prev = None
        cur = head
        for i in range(0, middle+1):
            prev = cur
            cur = cur.next

        prev.next = cur.next

        return head