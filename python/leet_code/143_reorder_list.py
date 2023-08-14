from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:

        list_len = 0
        cur = head
        while cur:
            list_len += 1
            cur = cur.next

        middle = list_len // 2
        index = 0

        prev = None
        cur = head
        while cur:
            if index < middle:
                cur = cur.next
            else:
                first_node = cur.next
                cur.next = prev
                prev = cur
                cur = first_node
            index += 1

        cur = head
        index = 0
        while index < middle:
            first_node = cur.next
            second_node = prev.next

            cur.next = prev
            prev.next = first_node
            cur = first_node
            prev = second_node

            index += 1
        cur.next = None

        return head


if __name__ == '__main__':
    sol = Solution()
    h = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5)))))
    sol.reorderList(h)
    # print_list(sol.reverseBetween(h, 1, 7))