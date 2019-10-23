#!/bin/python3
"""
You are given an unordered array consisting of consecutive integers 
$[1, 2, 3, ..., n]$ without any duplicates. You are allowed to swap any two
elements. You need to find the minimum number of swaps required to sort the
array in ascending order.

For example, given the array $arr = [7,1,3,2,4,5,6]$ we perform the following
steps:

i   arr                     swap (indices)
0   [7, 1, 3, 2, 4, 5, 6]   swap (0,3)
1   [2, 1, 3, 7, 4, 5, 6]   swap (0,1)
2   [1, 2, 3, 7, 4, 5, 6]   swap (3,4)
3   [1, 2, 3, 4, 7, 5, 6]   swap (4,5)
4   [1, 2, 3, 4, 5, 7, 6]   swap (5,6)
5   [1, 2, 3, 4, 5, 6, 7]
It took $5$ swaps to sort the array.

Approach: A* with distance from sorted as metrinc.

Take a priority queue, sorted by the edit distance and return
when the result is definitely (edit distance/2 + steps taken > cur).
Add all new elements to the priorityqueue [0].

[0]: https://golang.org/pkg/container/heap/
"""
import math
import os
import queue
import random
import re
import sys

from typing import List, Tuple

from dataclasses import dataclass, field
from typing import Any

@dataclass(order=True)
class PrioritizedItem:
    priority: int
    item: Tuple[int]=field(compare=False)
    wrong: frozenset=field(compare=False)
    edits: int


def distance_from_sorted(arr: List[int]) -> int:
    """
    How many elements in the array are in the wrong position?

    Since elements range from $0...n$, every element that is not
    where $x_i != i$ is in the wrong position.
    """
    distance = 0

    for idx, val in enumerate(arr):
        if idx + 1 != val:
            distance += 1

    return distance

# Complete the minimumSwaps function below.
def minimumSwaps(arr: List[int]):
    pq = queue.PriorityQueue()
    # Add first element to priority queue
    pq.put(PrioritizedItem(
        priority=distance_from_sorted(arr)/2,
        wrong=frozenset(idx for idx, val in enumerate(arr) if val != idx+1),
        item=tuple(arr),
        edits=0
    ))

    seen = set()

    while not pq.empty():
        item = pq.get()

        if not item.wrong:
            return item.edits

        # print(item)
        # for all positions in board, check the move
        for lhs in item.wrong:
            for rhs in item.wrong:
                if lhs==rhs:
                    continue

                # Not valid after swap -> a wasted swap
                if item.item[lhs] != rhs+1 and item.item[rhs] != lhs+1:
                    continue

                # Do the swap
                new_arr = list(item.item)
                new_arr[lhs] = item.item[rhs]
                new_arr[rhs] = item.item[lhs]

                # Seen?
                item_arr = tuple(new_arr)

                fixed = []
                if new_arr[lhs] == lhs+1:
                    fixed.append(lhs)
                if new_arr[rhs] == rhs+1:
                    fixed.append(rhs)

                if item_arr not in seen:
                    wrong = item.wrong - set(fixed)
                    pq.put(PrioritizedItem(
                        priority=item.edits+1+len(wrong)/2,
                        wrong=wrong,
                        item=item_arr,
                        edits=item.edits+1
                    ))

                    seen.add(item_arr)

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    res = minimumSwaps(arr)

    fptr.write(str(res) + '\n')

    fptr.close()
