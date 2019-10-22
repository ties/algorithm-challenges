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

#define PRIME_RANGE 65534
#define MAX_N 1120
#define MAX_K 14

#define UNKNOWN -1

static bool mark[PRIME_RANGE];

static int mem[MAX_N + 1][MAX_K + 1][MAX_N + 1];

inline int * idx(const int n, const int k, const int m) {
    return &mem[n][k][m];
}

/**
 * (#primes choose) | sum(..) == n
 * inductief:
 * (primes choose k-1) + i | sum(..) = n
 **/
int num_sets(const int N, const int K, const int M) {
    for (int n=1; n <= N; n++)
        for (int k=1; k <= K; k++)
			for (int m=1; m <= N; m++)
				*idx(n, k, m) = 0;
    
    // now init the k=1 positions for the primes;
    for (int n=2; n <= N; n++) {
        if (mark[n])
            *idx(n, 1, n) = 1;
    }

	for (int n=0; n <= N; n++) {
		for (int k=1; k <= K; k++) {
			// now: max from from max to M: only use "higher" primes
			for (int m=2; m <= N; n++) {
			}
		}
	}

	return -1;
}


int main(int argc, const char * argv[])
{
    int N, K;
    
    mark[0] = false;
    // init marked to true
    for (int i=2; i < PRIME_RANGE; i++)
        mark[i] = true;
    
    // calculate primes
    for (int i=2; i <= sqrt(PRIME_RANGE); i++) {
        if (mark[i]) {
            for (int j=i*i; j <= PRIME_RANGE; j += i) {
                mark[j] = false;
            }
        }
    }
    
    while (cin >> N >> K) {
        // count the different ways to sum N from K primes
        cout << num_sets(N, K, 0) << endl;
    }
    
    return 0;
}

