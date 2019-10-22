#include <iostream>
#include <stdio.h>
#include <math.h>
#include <algorithm>

#include <vector>
#include <list>
#include <map>
#include <set>

#include <queue>
#include <deque>
#include <stack>
#include <bitset>
#include <algorithm>
#include <functional>
#include <numeric>
#include <utility>
#include <sstream>
#include <iostream>
#include <fstream>
#include <iomanip>
#include <cstdio>
#include <cmath>
#include <cstdlib>
#include <ctime>
#include <cstring>
#include <assert.h>

using namespace std;

#define INVALID -1
#define UNKNOWN -2
#define MAX_SIZE 50+1

int N, K, M;

// memoization buffer |elements|, |bars|, len(cur)
// note that bar color (black/white) is iplicitly encoded in |bars|, since the first bar is black.
long mem[MAX_SIZE][MAX_SIZE][MAX_SIZE];

inline long* addr(const int n, const int k, const int m) {
	assert (n >= 1 && m >= 1 && k >= 1 && n <=N && m <= M && k <= K);
	return &mem[n][k][m];
}

// initialize the memoized data
void mem_init() {
	for (int n=1; n<=N; n++)
		for (int k=1; k<=K; k++)
			for (int m=1; m<=M; m++)
				*addr(n, k, m) = UNKNOWN;

	// add invalid initial solutions
	for (int k=1; k<=K; k++)
		for (int m=1; m<=M; m++)
			*addr(1, k, m) = INVALID;
	// add valid initial solution
	// (first bar is black, all other values for k/m are invalid)
	*addr(1,1,1) = 1;
}

long mem_count(const int n, const int k, const int m) {
	const long val = *addr(n, k, m);

	//printf("[memoize_retrieve] %d,%d,%d\n", n, k, m);
	
	if (val == INVALID)
		return 0;
	if (val != UNKNOWN)
		return val;

	// all n=1 values should be pre-memoized
	assert (n > 1);

	// calculate new value
	long newVal = 0;

	// create new bar
	// (n-1, k-1, any m) -> (n, k, 1)
	// iff k <= K and m==1
	if (m==1 && k <= K && k > 1) {
		//printf("[state] append-creating %d,%d,%d\n", n, k, m);
		for (int len=1; len <= M; len++) {
			//printf("[state] appending bar to %d,%d,%d\n", n-1, k-1, len);
			const long val = mem_count(n-1, k-1, len);

			if (val == INVALID)
				continue;

			newVal += val;
		}
	}

	// append to old bar
	// (n-1, k, m-1) -> (n, k, m)
	// iff m <= M
	if (m <= M && m > 1) {
		//printf("[state] appending to %d,%d,%d\n", n-1, k, m-1);
		const long val = mem_count(n-1, k, m-1);

		if (val != INVALID)
			newVal += val;
	}

	// store
	*addr(n, k, m) = newVal;
	return newVal;
}

int main(int argc, const char * argv[])
{
	// n = units, k = bars, m = max width
    while (cin >> N >> K >> M) {
		// reset memoization buffer
		mem_init();

		long res = 0;

		for (int m=1; m<=M; m++) {
			//printf("[lookup] %d,%d,%d\n", N, K, m);
			const long val = mem_count(N, K, m);

			if (val != INVALID)
				res += val;
		}

		cout << res << endl;
    }

    return 0;
}
