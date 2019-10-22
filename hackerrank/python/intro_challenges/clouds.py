#!/bin/python

import math
import os
import random
import re
import sys

# Complete the jumpingOnClouds function below.
def jumpingOnClouds(c):
    """
    Recursive solution: calculate number of jumps from current position
    to end of c.

    Params:
      c - int[]: array of binary integers, 0 is safe, 1 is not.
    """
    if len(c) == 1:
        return 0

    if len(c) == 2:
        # jump to next and done.
        return 1

    if c[2] == 0:
        return 1 + jumpingOnClouds(c[2:])
    else:
        assert c[1] == 0
        return 1 + jumpingOnClouds(c[1:])

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(raw_input())

    c = map(int, raw_input().rstrip().split())

    result = jumpingOnClouds(c)

    fptr.write(str(result) + '\n')

    fptr.close()
