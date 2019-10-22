#!/bin/python

import math
import os
import random
import re
import sys

# Complete the countingValleys function below.
def countingValleys(n, s):
    count = 0
    depth = 0
    for c in s:
        if c == 'D':
            depth -= 1
        elif c == 'U':
            if depth == -1:
                count += 1
            depth += 1
    return count


# print(countingValleys(8, 'UDDDUDUU'))
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(raw_input())

    s = raw_input()

    result = countingValleys(n, s)

    fptr.write(str(result) + '\n')

    fptr.close()
