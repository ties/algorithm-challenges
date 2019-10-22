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

struct barCode {
    int bars;
    int streak;
    bool last;
};
// bars, streak, last
typedef struct barCode bar_code;

static int N;
static int K;
static int M;

const inline bool is_valid(const int pos, bar_code mask) {
	if (mask.bars > K || (mask.bars + ((N-pos)/2) < K))
		return false;

	if (mask.streak > M)
		return false;

	return true;
}

const inline bar_code new_code(bar_code old, bool append) {
	bar_code res;

	res.last = append;

	if (old.last == append) {
		res.bars = old.bars;
		res.streak = old.streak + 1;
	} else {
		res.bars = old.bars + 1;
		res.streak = 1;
	}

	return res;
}

int main(int argc, const char * argv[])
{
    // n = units, k = bars, m = max width
    while (cin >> N >> K >> M) {
    	//printf("[init] n=%d b=%d streak=%d\n", N, K, M);

		vector<vector<bar_code> > prefixes(N+1);

		// start with dark (!) bar on the first position;
		bar_code b_start; 
		
		b_start.streak = 1;
		b_start.bars = 1;
		b_start.last = true;
		
		prefixes[0].push_back(b_start);

		for (int n=1; n< N; n++) {
			const vector<bar_code> prev = prefixes[n-1];
			vector<bar_code>& cur = prefixes[n];
			// generate possibilities based on shorter prefixes
			for (vector<bar_code>::const_iterator it = prev.begin(); it != prev.end(); ++it) {
				// generate the two possibilities
				bar_code b1 = new_code((*it), true);
				bar_code b2 = new_code((*it), false);

				if (is_valid(n, b1)) {
					cur.push_back(b1);
				}
				if (is_valid(n, b2)) {
					cur.push_back(b2);
				}
			}

			//printf("[n=%d] %d items\n", n, cur.size());
		}

		int solutions = 0;

		for (vector<bar_code>::iterator it = prefixes[N-1].begin(); it != prefixes[N-1].end(); ++it) {
			bar_code cur = *it;
			if (cur.bars == K) {
				//cout << "valid: " << *it << endl;
				solutions++;
			}
		}

    	//printf("[final] solutions %d\n", solutions);
		cout << solutions << endl;
    }

    return 0;
}
