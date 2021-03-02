#!/bin/python3

import math
import os
import random
import re
import sys

from collections import Counter

# Complete the isValid function below.
def isValid(s):
    char_counts = Counter(s)

    occurence_counts = Counter(char_counts.values())

    if len(occurence_counts) <= 1:
        return "YES"
    elif len(occurence_counts) > 2:
        return "NO"
    else:
        # Look at the second most common occurence
        (k_0, v_0), (k_1, v_1) = occurence_counts.most_common(2)

        # If the second most common occurence appears once, and is off by one
        # we can delete a char to fix it.
        if v_1 == 1 and k_1 - k_0 == 1:
            return "YES"
        return "NO"

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = isValid(s)

    fptr.write(result + '\n')

    fptr.close()

