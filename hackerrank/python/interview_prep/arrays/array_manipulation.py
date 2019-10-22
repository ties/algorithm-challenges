#!/bin/python3
"""
Starting with a 1-indexed array of zeros and a list of operations, for each
operation add a value to each of the array element between two given indices,
*inclusive*.

Once all operations have been performed, return the maximum value in your array.

For example, the length of your array of zeros $n=10$. Your list of queries is
as follows:
```
    a b k
    1 5 3
    4 8 7
    6 9 1
```
Add the values of $k$ between the indices $a$ and $b$ inclusive:
```
index->	 1 2 3  4  5 6 7 8 9 10
	[0,0,0, 0, 0,0,0,0,0, 0]
	[3,3,3, 3, 3,0,0,0,0, 0]
	[3,3,3,10,10,7,7,7,0, 0]
	[3,3,3,10,10,8,8,8,1, 0]
```

The largest value is $10$ after all operations are performed.

Solution
========
Work with integer intervals. But is too slow for large inputs (oops).
"""
import bisect
import itertools
import math
import os
import random
import re
import sys

from typing import List, NamedTuple

class Query(NamedTuple):
    lhs: int
    rhs: int
    k: int

class Slice:
    lhs: int = 0
    rhs: int = 0
    val: int = 0

    def __init__(self, lhs: int, rhs: int, val: int) -> None:
        self.lhs = lhs
        self.rhs = rhs
        self.val = val

    def overlaps(self, q: Query) -> bool:
        # Check for complete coverage:
        if q.lhs <=self.lhs and q.rhs >= self.rhs:
            return True
        # Check for containment, q ends before self
        if q.lhs <= self.rhs and q.rhs >= self.lhs and q.rhs <= self.rhs:
            # Starts before our end, it overlaps, but ends before
            return True
        if q.lhs >= self.lhs and q.lhs <= self.rhs:
            return True

        return False

    def add(self, q: Query) -> List['Slice']:
        # Does not overlap -> throw
        if not self.overlaps(q):
            return []
        # Base case: complete overlap, change own value
        if q.lhs <= self.lhs and q.rhs >= self.rhs:
            return [Slice(self.lhs, self.rhs, self.val + q.k)]

        parts = []
        if q.lhs <= self.lhs:
            # Starts at left boundary
            # Make part up to rhs
            right_bound = self.rhs if self.rhs < q.rhs else q.rhs
            parts.append(Slice(self.lhs, right_bound, self.val + q.k))
            # Make part after query
            if q.rhs < self.rhs:
                parts.append(Slice(q.rhs+1, self.rhs, self.val))
        else:
            # Query starts in this slice.
            parts.append(Slice(self.lhs, q.lhs-1, self.val))
            # Now start at left boundary of query
            right_bound = self.rhs if self.rhs < q.rhs else q.rhs
            parts.append(Slice(q.lhs, right_bound, self.val + q.k))
            # Make part after query
            if q.rhs < self.rhs:
                parts.append(Slice(q.rhs+1, self.rhs, self.val))

        return parts

    def __repr__(self) -> str:
        return f"Slice(lhs={self.lhs}, rhs={self.rhs}, val={self.val})"


class SliceSet:
    lhs: List[int] = []
    rhs: List[int] = []
    val: List[int] = []

    def __init__(self, slices: List[Slice]) -> None:
        for slice in slices:
            self.lhs.append(slice.lhs)
            self.rhs.append(slice.rhs)
            self.val.append(slice.val)

    def add(self, q: Query) -> None:
        # print(f"state: {len(self.lhs)} elements")
        # Find the position q applies to
        start_idx = bisect.bisect_left(self.rhs, q.lhs)
        end_idx = bisect.bisect_right(self.lhs, q.rhs)

        # Restrict to full overlap.
        if end_idx < start_idx:
            end_idx = start_idx

        results = []

        for idx in range(start_idx, end_idx):
            slice = Slice(self.lhs[idx], self.rhs[idx], self.val[idx])
            results.extend(slice.add(q))

        self.lhs = self.lhs[:start_idx] + [s.lhs for s in results] + self.lhs[end_idx:]
        self.rhs = self.rhs[:start_idx] + [s.rhs for s in results] + self.rhs[end_idx:]
        self.val = self.val[:start_idx] + [s.val for s in results] + self.val[end_idx:]

        # import ipdb
        # ipdb.set_trace()

# Complete the arrayManipulation function below.
def arrayManipulation(n, queries):
    """
    Ignoring the potential merging of slices that have same value.
    """
    # Array of integer slices
    slices = SliceSet([Slice(1, n, 0)])

    for lhs, rhs, k in queries:
        if not lhs <= rhs:
            continue
        slices.add(Query(lhs, rhs, k))

    # Find maximum value
    return max(slices.val)

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    nm = input().split()

    n = int(nm[0])

    m = int(nm[1])

    queries = []

    for _ in range(m):
        queries.append(list(map(int, input().rstrip().split())))

    result = arrayManipulation(n, queries)

    fptr.write(str(result) + '\n')

    fptr.close()
