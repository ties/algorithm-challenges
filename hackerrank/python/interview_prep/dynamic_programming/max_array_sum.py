
#!/bin/python

import math
import os
import random
import re
import sys

# Complete the maxSubsetSum function below.
def maxSubsetSum(arr):
    """
    Calculate the maximum possible sum of non-adjacent elements.
    """
    n = len(arr)
    subsum_included = [0]*n
    subsum_not_included = [0]*n

    subsum_not_included[n-1] = 0
    subsum_included[n-1] = arr[n-1]

    for i in range(n-2, -1, -1):
        subsum_included[i] = arr[i] + subsum_not_included[i+1]
        subsum_not_included[i] = max(
            subsum_not_included[i+1],
            subsum_included[i+1]
        )
        # print(i)
        # print(subsum_included)
        # print(subsum_not_included)
        # print("===")

    return max(subsum_included[0], subsum_not_included[0])



if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(raw_input())
    arr = map(int, raw_input().rstrip().split())

    res = maxSubsetSum(arr)

    fptr.write(str(res) + '\n')
    fptr.close()
