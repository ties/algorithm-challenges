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

static int mem[MAX_N + 1][MAX_K + 1];

typedef pair<int, int> int_t2;

inline int * idx(const int n, const int k) {
    return &mem[n][k];
}

/**
 * (#primes choose) | sum(..) == n
 * inductief:
 * (primes choose k-1) + i | sum(..) = n
 **/
int num_sets(const int N, const int K) {
    //printf("[mem] num_sets(%d, %d)\n", N, K);
    if (*idx(N, K) != UNKNOWN)
        return *idx(N, K);
    
    int total = 0;
    
	set<int_t2> seen;
    // gebruik volgorde om niet dubbel te tellen
    for (int n=2; n < N; n++) {
        if (mark[n] && N-n >= 0 && !seen.count(int_t2(n, N-n))) {
			total += num_sets(N - n, K-1);

			seen.insert(int_t2(N-n, n));
			seen.insert(int_t2(n, N-n));
            
            if (num_sets(N-n, K-1) > 0)
                printf("[count] %d + %d = %d (%dx)\n", N-n, n, N, num_sets(N-n, K-1));
        }
    }
    
    *idx(N, K) = total;
    
    return total;
}


/*
 reset memory
 */
void mem_reset(const int N, const int K) {
    for (int n=0; n <= N; n++)
        for (int k=0; k <= K; k++)
            *idx(n, k) = UNKNOWN;
    
    // now init the k=1 positions for the primes;
    for (int n=2; n <= N; n++) {
        if (mark[n])
            *idx(n, 1) = 1;
        else
            *idx(n, 1) = 0;
    }
    
    // init the (0, 0) value.
    *idx(0, 1) = 0;
	*idx(1, 1) = 0;
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
            for (int j=i*i; j < PRIME_RANGE; j += i) {
                mark[j] = false;
            }
        }
    }
    
    // now copy them into the vector
    int primeCount = 0;
    for (int i=0; i < PRIME_RANGE; i++) {
        if (mark[i]) {
            primeCount++;
        }
    }
    
    //printf("[primes] %d primes found\n", primeCount);
    
    while (cin >> N >> K) {
        // reset the memoization buffer
        mem_reset(N, K);
        // count the different ways to sum N from K primes
        cout << num_sets(N, K) << endl;
    }
    
    return 0;
}

