import unittest
from typing import Optional, List

if __name__ == "__main__":
    unittest.main()

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def longest_balanced_parentheses(self, s: str) -> int:
    stack = [-1]
    max_length = 0

    for i, char in enumerate(s):
        if char == '(':
            stack.append(i)
        else:
            stack.pop()

            if not stack:
                stack.append(i)
            else:
                max_length = max(max_length, i - stack[-1])
    return max_length

def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
    rev = None
    while head is not None:
        temp = head.next
        head.next = rev
        rev = head
        head = temp
    return rev

class ConstantTimeSet:
    def __init__(self, n: int):
        self.n = n
        self.m = [''] * (n + 1)

    def insert(self, k: int, v: str):
        if 1 <= k <= self.n:
            self.m[k] = v

    def delete(self, k: int):
        if 1 <= k <= self.n:
            self.m[k] = ''

    def search(self, k: int) -> bool:
        if 1 <= k <= self.n:
            return self.m[k] == ''
        return False

def data_overhead(data_space: int, total_space: int) -> float:
    return 1 - (data_space/total_space)

import sys

class SegmentTree:
    def __init__(self, arr: List[int]):
        self.n = len(arr)
        self.tree = [0] * (4 * self.n)
        self._build(arr, 0, 0, self.n - 1)

    def _build(self, arr, node, start, end):
        if start == end:
            self.tree[node] = arr[start]
            return
        mid = (start + end) // 2
        left_child = 2 * node + 1
        right_child = 2 * node + 2
        self._build(arr, left_child, start, mid)
        self._build(arr, right_child, mid + 1, end)
        self.tree[node] = min(self.tree[left_child], self.tree[right_child])

    def _query(self, node, start, end, i, j):
        if i > end or j < start:
            return sys.maxsize
        if i <= start and end <= j:
            return self.tree[node]
        # Partial overlap
        mid = (start + end) // 2
        left_min = self._query(2 * node + 1, start, mid, i, j)
        right_min = self._query(2 * node + 2, mid + 1, end, i, j)
        return min(left_min, right_min)

    def range_min(self, i, j):
        return self._query(0, 0, self.n - 1, i, j)

    def _update(self, node, start, end, idx, new_val):
        if start == end:
            self.tree[node] = new_val
        else:
            mid = (start + end) // 2
            if idx <= mid:
                self._update(2 * node + 1, start, mid, idx, new_val)
            else:
                self._update(2 * node + 2, mid + 1, end, idx, new_val)
            self.tree[node] = min(self.tree[2 * node + 1], self.tree[2 * node + 2])

    def update(self, idx, new_val):
        self._update(0, 0, self.n - 1, idx, new_val)

class SegmentTree:
    def __init__(self, xs: List[int]):
        self.xs = xs
        self.n = len(xs) - 1
        self.count = [0] * (4 * self.n)
        self.covered = [0] * (4 * self.n)

    # be careful when you are reading over, we use right + 1 instead of right.
    # make sure you understand why.
    def update(self, qleft, qright, qval, left, right, pos):
        if self.xs[right+1] <= qleft or self.xs[left] >= qright:
            return
        if qleft <= self.xs[left] and self.xs[right+1] <= qright: # full overlap
            self.count[pos] += qval
        else: # partial overlap
            mid = (left + right) // 2
            self.update(qleft, qright, qval, left, mid, pos*2 + 1)
            self.update(qleft, qright, qval, mid+1, right, pos*2 + 2)

        if self.count[pos] > 0:
            self.covered[pos] = self.xs[right+1] - self.xs[left]
        else:
            if left == right:
                self.covered[pos] = 0
            else:
                self.covered[pos] = self.covered[pos*2 + 1] + self.covered[pos*2 + 2]

    def query(self):
        return self.covered[0]


class TicTacToe:
    def __init__(self, n):
        self.n = n
        self.rows = [0] * n
        self.cols = [0] * n
        self.diag = 0
        self.anti_diag = 0

    def move(self, row, col, player):
        # Player 1: +1, Player 2: -1
        to_add = 1 if player == 1 else -1

        self.rows[row] += to_add
        self.cols[col] += to_add
        if row == col:
            self.diag += to_add
        if row + col == self.n - 1:
            self.anti_diag += to_add

        # Check win
        if (abs(self.rows[row]) == self.n or
                abs(self.cols[col]) == self.n or
                abs(self.diag) == self.n or
                abs(self.anti_diag) == self.n):
            return player

        return 0

class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

def bst_to_dll(root):
    last, head = None, None

    def helper(node):
        nonlocal last, head
        if not node:
            return

        helper(node.left)
        if last:
            last.right = node
            node.left = last
        else:
            head = node
        last = node
        helper(node.right)


    helper(root)
    return head

def merge_dlls(l1, l2):
    dummy = TreeNode(0)
    curr = dummy

    while l1 and l2:
        if l1.val < l2.val:
            curr.right = l1
            l1.left = curr
            l1 = l1.right
        else:
            curr.right = l2
            l2.left = curr
            l2 = l2.right
        curr = curr.right

    remaining = l1 if l1 else l2
    while remaining:
        curr.right = remaining
        remaining.left = curr
        curr = curr.right
        remaining = remaining.right

    head = dummy.right
    if head:
        head.left = None
    return head

def balance_bst(root):
    def inorder_traversal(node, result):
        if node:
            inorder_traversal(node.left, result)
            result.append(node.val)
            inorder_traversal(node.right, result)

    def helper(a, l, r):
        if l > r:
            return None

        mid = (l+r)//2
        node = a[mid]
        node.left = helper(a, l, mid-1)
        node.right = helper(a, mid+1, r)
        return node

    values = []
    inorder_traversal(root, values)
    return helper(values, 0, len(values) - 1)

def is_height_balanced(root):
    def helper(node):
        if not node:
            return 0, True

        left_height, left_balanced = helper(node.left)
        right_height, right_balanced = helper(node.right)

        height = max(left_height, right_height) + 1
        balanced = (
                left_balanced and
                right_balanced and
                abs(left_height - right_height) <= 1
        )

        return height, balanced

    _, is_bal = helper(root)
    return is_bal

def rev_ll(head):
    rev = None

    while head is not None:
        temp = head.next
        head.next = rev
        rev = head
        head = temp
    return rev

def ch3p44(a):
    n = len(a)
    m1 = [1] * n
    m2 = [1] * n

    for i in range(1, n):
        m1[i] = m1[i-1] * a[i-1]
    for i in range(n-2, -1, -1):
        m2[i] = m2[i+1] * a[i+1]

    res = [0] * n
    for i in range(n):
        res[i] = m1[i] * m2[i]

    return res


class Tests(unittest.TestCase):

    # Example Usage
    arr = [1, 3, 2, 7, 9, 11, 4, 6]
    st = SegmentTree(arr)

    print("Minimum in range [1, 4]:", st.range_min(1, 4))  # Output: 2
    print("Minimum in range [3, 7]:", st.range_min(3, 7))  # Output: 4


    def test_longest_balanced_parentheses(self):
        test_cases = [
            (")()(())()()))())))(", 10),
            (")(", 0),
            ("(())", 4),
        ]

        for s, expected in test_cases:
            with self.subTest(s=s):
                v = longest_balanced_parentheses(s)
                print("VALUE",v )
                self.assertEqual(v, expected)
