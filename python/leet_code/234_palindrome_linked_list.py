from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        list_len = 0

        cur = head
        while cur:
            list_len += 1
            cur = cur.next

        if list_len // 2 == 0:
            return True

        prev = None
        first_head = head

        index = 0
        while index < (list_len // 2):
            index += 1
            next_node = first_head.next
            first_head.next = prev
            prev = first_head
            first_head = next_node

        if list_len % 2:
            first_head = first_head.next

        index = 0
        while index < (list_len // 2):
            index += 1
            if first_head.val != prev.val:
                return False
            first_head = first_head.next
            prev = prev.next

        return True


if __name__ == '__main__':
    sol = Solution()
    h = ListNode(1, ListNode(2, ListNode(1)))
    sol.isPalindrome(h)
    # print_list(sol.reverseBetween(h, 1, 7))