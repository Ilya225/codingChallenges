from typing import Optional


class ListNode:
    def __init__(self, x, next = None):
        self.val = x
        self.next = next

class Solution:
    def reverseList_(self, head: ListNode) -> ListNode:
        if head == None or head.next == None:
            return head

        child = self.reverseList_(head.next)
        head.next.next = head
        head.next = None
        # child.next.next = None
        return child

    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:

        prev = None
        cur = head

        while cur:
            next_node = cur.next
            cur.next = prev
            prev = cur
            cur = next_node

        head = prev

        return head



if __name__ == '__main__':
    sol = Solution()
    h = ListNode(1, ListNode(2, ListNode(3)))
    sol.reverseList_(h)
