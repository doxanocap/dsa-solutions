from typing import List, Union
import heapq

class SegmentTree:
    def __init__(self, xs: List[int]):
        self.xs = xs
        self.n = len(xs)
        self.count = [0] * (4 * self.n)
        self.covered = [0] * (4 * self.n)

    def _update(self, ql, qr, qv, l, r, idx):
        if self.xs[r] <= l or r <= self.xs[l]:
            return
        if ql <= self.xs[l]  and self.xs[r] <= qr:
            self.count[idx] += 1
        else:
            m = l + (r-l)//2
            self._update(ql,qr,qv,l,m, idx*2+1)
            self._update(ql,qr,qv,m+1,r, idx*2+2)

        if self.count[idx] > 0:
            self.covered[idx] = self.xs[r] - self.xs[l]
        else:
            self.covered[idx] = self.covered[idx+2+1] + self.covered[idx*2+2]

    def update(self, ql, qr, qv):
        return self._update(ql, qr, qv, 0, self.n-1, 0)

    def query(self):
        return self.covered[0]

class Solution:
    def separateSquares(self, squares: List[List[int]]) -> float:
        events = []
        xs_set = set()
        for x, y, l in squares:
            events.append((y, 1, x, x+l))
            events.append((y+l, -1, x, x+l))
            xs_set.update([x, x + l])
        xs = sorted(xs_set)

        stree = SegmentTree(xs)
        events.sort()

        total_area = 0.0
        prev_y = events[0][0]
        for y, start, xl, xr in events:
            curr_width = stree.query()
            total_area += curr_width * (y - prev_y)
            stree.update(xl, xr, start)
            prev_y = y

        stree = SegmentTree(xs)
        curr_area = 0.0
        prev_y = events[0][0]
        for y, start, xl, xr in events:
            curr_width = stree.query()
            if curr_area + curr_width * (y - prev_y) >= total_area / 2.0:
                return prev_y + (total_area / 2.0 - curr_area) / curr_width
            curr_area += curr_width * (y - prev_y)
            stree.update(xl, xr, 0)
            prev_y = y