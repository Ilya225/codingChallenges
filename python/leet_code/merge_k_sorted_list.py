from typing import List
from queue import PriorityQueue


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:
        root = None
        nextLink = None

        stopCounter=0
        minIndex=None

        if not len(lists):
            return root

        while stopCounter < len(lists):
            for i, p in enumerate(lists):
                if p is None:
                    continue

                if minIndex is None:
                    minIndex = i

                if p.val < lists[minIndex].val:
                    minIndex = i
            if minIndex is None:
                break

            if root is None:
                root = ListNode(lists[minIndex].val)
                nextLink = root
            else:
                nextLink.next = ListNode(lists[minIndex].val)
                nextLink = nextLink.next


            if lists[minIndex].next is None:
                stopCounter += 1

            lists[minIndex] = lists[minIndex].next
            minIndex = None

        return root

    def mergeQueue(self, lists :List[ListNode]) -> ListNode:
        root = None
        kmins = PriorityQueue()

        for i,l in enumerate(lists):
            if l is None: continue
            kmins.put((l.val, i))

        if kmins.empty():
            return root

        entry = kmins.get()
        root = ListNode(entry[0])
        nextLink = root
        if lists[entry[1]].next is not None:
            lists[entry[1]] = lists[entry[1]].next
            kmins.put((lists[entry[1]].val, entry[1]))


        while not kmins.empty():
            entry = kmins.get()
            nextLink.next = ListNode(entry[0])
            nextLink = nextLink.next
            if lists[entry[1]].next is not None:
                lists[entry[1]] = lists[entry[1]].next
                kmins.put((lists[entry[1]].val, entry[1]))

        return root




if __name__ == '__main__':
    sol = Solution()
    sol.mergeKLists([])