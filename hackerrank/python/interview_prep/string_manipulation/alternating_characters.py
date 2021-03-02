"""
You are given a string containing characters  and  only. Your task is to change
it into a string such that there are no matching adjacent characters. To do this,
you are allowed to delete zero or more characters in the string.

Your task is to find the minimum number of required deletions.

## Example
$$s = AABAAB$$

Remove an $A$ at positions $0$ and $3$ to make  in  deletions.

## Function Description

Complete the alternatingCharacters function in the editor below.

alternatingCharacters has the following parameter(s):
  * string s: a string

### Returns
  * int: the minimum number of deletions required

### Input Format

The first line contains an integer $q$, the number of queries.
The next $q$ lines each contain a string $s$ to analyze.

Constraints
  * $1 <= q <= 10$
  * $1 <= length of s <= 10^5$
  * Each string $s$ will consist only of characters $A$ and $B$.

### Sample Input

```
5
AAAA
BBBBB
ABABABAB
BABABA
AAABBB
```

### Sample Output
```
3
4
0
0
4
```

### Explanation

The characters marked red are the ones that can be deleted so that the string
does not have matching adjacent characters.
"""
#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the alternatingCharacters function below.
def alternatingCharacters(s):
    deletes = 0
    for i in range(0, len(s) - 1):
        if s[i] == s[i+1]:
            deletes += 1

    return deletes



def alternatingCharacters_recursive(s):
    # Base case
    if len(s) <= 1:
        return 0

    if s[0] != s[1]:
        return alternatingCharacters(s[1:])
    else:
        return 1 + alternatingCharacters(s[1:])


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input())

    for q_itr in range(q):
        s = input()

        result = alternatingCharacters(s)

        fptr.write(str(result) + '\n')

    fptr.close()
