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


def swap(arr: List[int], lookup: List[int], lhs: int, rhs: int) -> None:
    """Directly swap two items."""
    # print(f"swap({arr}, {lookup}, {lhs}, {rhs})")
    # First update lookup
    lookup[arr[rhs]] = lhs
    lookup[arr[lhs]] = rhs

    # Afterwards, swap
    tmp = arr[lhs]
    arr[lhs] = arr[rhs]
    arr[rhs] = tmp

# Complete the minimumSwaps function below.
def minimumSwaps(arr: List[int]):
    lookup = [0] * (len(arr)+1)
    for idx, val in enumerate(arr):
        lookup[val] = idx

    edits = 0

    for i in range(len(arr)):
        if arr[i] != i+1:
            edits += 1
            swap(arr, lookup, i, lookup[i+1])

    return edits



if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    res = minimumSwaps(arr)

    fptr.write(str(res) + '\n')

    fptr.close()
