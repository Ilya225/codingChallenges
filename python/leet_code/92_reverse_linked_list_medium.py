from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def print_list(h: ListNode):
    res = "["

    while h:
        res += " " + str(h.val)
        h = h.next
        if h:
            res += ","

    res += "]"

    print(res)
    return res


class Solution:

    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:

        start = None
        cur = head
        count = 1

        while cur and count != left:
            count += 1
            start = cur
            cur = cur.next

        if not cur:
            return head

        prev = None
        tmp_head = cur

        while cur and count != right:
            count += 1
            next_node = cur.next
            cur.next = prev
            prev = cur
            cur = next_node

        if cur:
            next_node = cur.next
            cur.next = prev
            prev = cur
            cur = next_node

        if start is not None:
            start.next = prev
        else:
            head = prev

        if cur:
            tmp_head.next = cur


        return head





if __name__ == '__main__':
    sol = Solution()
    h = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5, ListNode(6, ListNode(7)))))))
    print_list(sol.reverseBetween(h, 1, 7))


