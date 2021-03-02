#!/bin/python3

import math
import os
import random
import re
import sys

def drop_middle(s: str) -> str:
    len_s = len(s)
    if len_s % 2 == 0:
        return s
    if len_s < 3:
        return s

    half_len = len_s // 2
    return s[0:half_len] + s[half_len+1:]

def all_the_same(s: str) -> bool:
    len_s = len(s)
    midpoint = -1
    if len_s >= 3 and len_s % 2 == 1:
        midpoint = len_s // 2 + 1

    for i in range(1, len(s)):
        if i != midpoint and s[i] != s[0]:
            return False
    return True


# Complete the substrCount function below.
def substrCount(n, s):
    num_special_substr = 0

    for i in range(n):
        for j in range(i+1, n+1):
            s_test = drop_middle(s[i:j])
            if all_the_same(s_test):
                num_special_substr += 1

    return num_special_substr


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    s = input()

    result = substrCount(n, s)

    fptr.write(str(result) + '\n')

    fptr.close()
