import collections, fileinput
from math import floor, factorial as fact

"""
1-based indexing is used in this file, because the index is relevant in some calculations.

The input is a line, consisting of multiple letters
each letter should be used once

The ordering of the lines is alphebetical 

**example of ordering**:
```
abaabc
abaacb
ababac
```

**example input:**
```
abaabc
cbbaa
#
```
**example output:**
```
ababac
No Successor
````

## Analysis
Since every character should be used in the output, this output is not
the successor in the lexicalcographical order of subsets.

However the successor likely is the next lexicalcographical permutation.

**Approach**
  * Introduce an order on the equivalent characters
  * `Rank(s)`
  * `UnRank(Rank(s)+1)`

or just use the lexicalcographical permutation successor algorithm.
"""

"""
Rank the given permutation

Kreher and Stinson, Algorithm 2.15 PermLexRank(n, pi)
"""
def rank(n, pi):
	r = 0
	rho = list(pi)

	for j in range(1, n+1):
		r = r + (rho[j] - 1) * fact(n-j)
		for i in range(j+1, n+1):
			if rho[i] > rho[j]:
				rho[i] = rho[i] - 1

	return r


"""
Unrank the permutation with rank *m* of length *n*

Kreher and Stinson, Algorithm 2.15 PermLexUnRank(n, r)
"""
def unrank(n, r, pi):
	pi[n] = 1
	for j in range(1, n):
		# floor, otherwise the division has a floating point remainder.0
		d = floor((r % fact(j+1))/fact(j))
		r = r - d * fact(j)
		pi[n-j] = d+1
		for i in range(n-j+1, n+1):
			if pi[i] > d:
				pi[i] = pi[i] + 1
	return pi

"""
Prepare the permuation/subsituttion arrays with 1-based indexes.
"""
def lexicographical_successor(letters):
	ordered_letters = [0] + sorted(letters)

	# prepare the character set
	# transform (aba -> a_1, b_1, a_2)
	tmp = list(ordered_letters)
	ordinals = [0]
	for c in letters:
		idx = tmp.index(c)
		tmp[idx] = None
		ordinals.append(idx)

	# rank the string of ordinals
	r = rank(len(letters), ordinals)

	# return the successor if it exists.
	u = unrank(len(letters), r+1, ordinals)

	# subsitute:
	as_letters = ''.join(list(map(lambda i: ordered_letters[i], u[1:])))

	if as_letters == letters:
		return 'No Successor'

	return as_letters

# read the input lines up to the hash
for line in map(lambda s: s.strip(), fileinput.input()):
	if line == '#':
		break

	print(lexicographical_successor(line))
