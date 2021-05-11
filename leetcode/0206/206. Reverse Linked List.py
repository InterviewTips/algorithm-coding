# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        pre = None
        cur = head
        while cur:
            next_node = cur.next  # 记录 next
            cur.next = pre  # cur 指向 pre
            pre = cur  # pre 前移
            cur = next_node  # cur 前移
        return pre
