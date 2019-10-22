#!/bin/python

import math
import os
import random
import re
import sys

def num_of_char(s, c_match):
    match = 0
    for c in s:
        if c == c_match:
            match += 1

    return match


# Complete the repeatedString function below.
def repeatedString(s, n):
    """Count the occurences of a in the string `s` repeated until it is 'n'
    chars long."""
    # Number of compleet repetitions.
    complete = n // len(s)
    prefix_len = n % len(s)

    return num_of_char(s, 'a')*complete + num_of_char(s[:prefix_len], 'a')

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = raw_input()

    n = int(raw_input())

    result = repeatedString(s, n)

    fptr.write(str(result) + '\n')

    fptr.close()
