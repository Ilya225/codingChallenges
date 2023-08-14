# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


from typing import Optional


class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:

        result = None
        root = None
        while list1 is not None and list2 is not None:
            if list1.val < list2.val:
                new_node = ListNode(list1.val)
                list1 = list1.next
            else:
                new_node = ListNode(list2.val)
                list2 = list2.next

            if root is None:
                result = new_node
                root = new_node
            else:
                result.next = new_node
                result = result.next

        last_list = list1 if list1 is not None else list2

        while last_list is not None:
            new_node = ListNode(last_list.val)
            last_list = last_list.next

            if root is None:
                result = new_node
                root = new_node
            else:
                result.next = new_node
                result = result.next

        return root
