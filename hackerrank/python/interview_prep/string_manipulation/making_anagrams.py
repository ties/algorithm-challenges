#!/bin/python3

import math
import os
import random
import re
import sys

from collections import Counter
# Complete the makeAnagram function below.
def makeAnagram(a, b):
    c_a = Counter(a)
    c_b = Counter(b)

    i = 0;

    all_keys = frozenset(c_a.keys()) | frozenset(c_b.keys())

    for k in all_keys:
        a = c_a.get(k, 0)
        b = c_b.get(k, 0)

        if a > b:
            i += a-b
        elif a < b:
            i += b-a

    return i

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    a = input()

    b = input()

    res = makeAnagram(a, b)

    fptr.write(str(res) + '\n')

    fptr.close()

