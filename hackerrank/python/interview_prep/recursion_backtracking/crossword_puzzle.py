#!/bin/python3
"""
A $10x10$ Crossword grid is provided to you, along with a set of words (or
names of places) which need to be filled into the grid. Cells are marked either
+ or -. Cells marked with a - are to be filled with the word list.

The following shows an example crossword from the input $crossover$ grid and
the list of words to fit, $words$:

```
Input 	   		Output

++++++++++ 		++++++++++
+------+++ 		+POLAND+++
+++-++++++ 		+++H++++++
+++-++++++ 		+++A++++++
+++-----++ 		+++SPAIN++
+++-++-+++ 		+++A++N+++
++++++-+++ 		++++++D+++
++++++-+++ 		++++++I+++
++++++-+++ 		++++++A+++
++++++++++ 		++++++++++
POLAND;LHASA;SPAIN;INDIA
```

Function Description
--------------------

Complete the crosswordPuzzle function in the editor below. It should return an
array of strings, each representing a row of the finished puzzle.

crosswordPuzzle has the following parameter(s):
    crossword: an array of  strings of length  representing the empty grid
    words: a string consisting of semicolon delimited strings to fit into
"""

import math
import os
import random
import re
import sys

# from dataclasses import dataclass
from enum import Enum
from typing import List, NamedTuple, Optional, FrozenSet


class Direction(Enum):
    HORIZONTAL = 'horizontal'
    VERTICAL = 'vertical'


class CrosswordSlot(NamedTuple):
    """A slot in a crossword puzzle"""
    direction: Direction
    i: int
    j: int
    length: int


class Crossword:
    crossword: List[List[str]]
    slots: FrozenSet[CrosswordSlot]

    def __init__(self, crossword: List[str],
                 slots: Optional[FrozenSet[CrosswordSlot]]=None) -> None:
        self.crossword = [list(row) for row in crossword]
        self.slots = self.__find_slots(crossword) if not slots else slots

    def __find_slots(self, crossword: List[str]) -> FrozenSet[CrosswordSlot]:
        """Find available slots."""
        slots: List[CrosswordSlot] = []
        for i in range(10):
            for j in range(10):
                # Horizontal
                if (((crossword[i][j] == '-') and
                     (j == 0 or crossword[i][j-1] != '-'))):
                    # Scan range
                    for end in range(j,11):
                        if end == 10:
                            break
                        if crossword[i][end] !='-':
                            break

                    if end-j > 1:
                        slots.append(CrosswordSlot(
                            direction=Direction.HORIZONTAL,
                            i=i,
                            j=j,
                            length=end-j
                        ))
                # Vertical
                if (((crossword[i][j] == '-' and
                     (i == 0 or crossword[i-1][j] != '-')))):
                    # Scan range
                    for end in range(i, 11):
                        if end == 10:
                            break
                        if crossword[end][j] != '-':
                            break

                    if end-i > 1:
                        slots.append(CrosswordSlot(
                            direction=Direction.VERTICAL,
                            i=i,
                            j=j,
                            length=end-i
                        ))
        return frozenset(slots)

    @property
    def empty_slots(self) -> List[CrosswordSlot]:
        # For every slot, check if they are completely not '-'
        empty = []
        for slot in self.slots:
            if not all(map(lambda s: s != '-', self.read_slot(slot))):
                empty.append(slot)

        return empty

    def read_slot(self, slot: CrosswordSlot) -> str:
        if slot.direction == Direction.HORIZONTAL:
            return ''.join(self.crossword[slot.i][slot.j:slot.j+slot.length])
        elif slot.direction == Direction.VERTICAL:
            out = []
            j = slot.j
            for i in range(slot.i, slot.i + slot.length):
                out.append(self.crossword[i][j])

            return ''.join(out)
        else:
            raise ValueError("Unknown direction")

    def place(self, slot: CrosswordSlot, word: str) -> 'Crossword':
        if slot.length != len(word):
            raise ValueError("Length mismatch")

        c = Crossword(self.crossword, self.slots)

        if slot.direction == Direction.HORIZONTAL:
            i = slot.i
            for idx, j in enumerate(range(slot.j, slot.j + slot.length)):
                val = c.crossword[i][j]
                if (val != '-' and val != word[idx]):
                    raise ValueError(f"{val} != {word[idx]} at [{i},{j}]")
            for idx, j in enumerate(range(slot.j, slot.j + slot.length)):
                c.crossword[i][j] = word[idx]
        elif slot.direction == Direction.VERTICAL:
            j = slot.j
            for idx, i in enumerate(range(slot.i, slot.i + slot.length)):
                val = c.crossword[i][j]
                if (val != '-' and val != word[idx]):
                    raise ValueError(f"{val} != {word[idx]} at [{i},{j}]")
            for idx, i in enumerate(range(slot.i, slot.i + slot.length)):
                c.crossword[i][j] = word[idx]

        return c

    def __repr__(self) -> str:
        return '\n'.join(''.join(row) for row in self.crossword)


def crossword_puzzle(cp: Crossword, words: FrozenSet[str]):
    if not words or not cp.empty_slots:
        return cp

    candidates: List[Crossword] = []
    slots = cp.empty_slots
    for word in words:
        for slot in slots:
            if len(word) == slot.length:
                try:
                    after = cp.place(slot, word)
                    tmp = crossword_puzzle(after, words - set([word]))

                    candidates.append(tmp)
                except ValueError:
                    pass

    # Sort by number of empty slots
    best = cp
    best_empty = len(cp.empty_slots)
    for c in candidates:
        empty = len(c.empty_slots)
        if empty < best_empty:
            best = c
            best_empty = empty

    return best

# Complete the crosswordPuzzle function below.
def crosswordPuzzle(crossword: List[str], words: List[str]):
    cp = Crossword(crossword)

    res = crossword_puzzle(cp, frozenset(words.split(';')))
    return str(res).split('\n')

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    crossword = []

    for _ in range(10):
        crossword_item = input()
        crossword.append(crossword_item)

    words = input()

    result = crosswordPuzzle(crossword, words)

    fptr.write('\n'.join(result))
    fptr.write('\n')

    fptr.close()
