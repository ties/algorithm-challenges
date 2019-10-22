#include <iostream>
#include <bitset>

using namespace std;

#define MIN(s1, s2) (s1 < s2 ? s1 : s2)
#define MAX(s1, s2) (s1 > s2 ? s1 : s2)

int cycleLength(unsigned int input);

/**
 * UVA "3n+1 problem" <http://uva.onlinejudge.org/external/1/100.pdf>
 */
int main(int argc, const char * argv[])
{
    unsigned int i1, i2;

    while (cin >> i1 >> i2) {
    	int max = 0;
    	for(int i=MIN(i1, i2); i <= MAX(i1, i2); i++) {
    		const int cl_i = cycleLength(i);
    		max = MAX(max, cl_i);
    	}

    	cout << i1 << " " << i2 << " " << max << endl;
    }
    return 0;
}

/**
 * The "3n+1" algorithm, counting variant.
 */
inline int cycleLength(unsigned int input) {
	unsigned int c = 1;

	unsigned int n = input;

	while(n != 1) {
		if (n % 2 == 1) { // mask the final bit & see if it's odd (vs %)
			n = 3*n+1;
		} else {
			// multiple iterations possible at once if there are trailing ones:
			unsigned int trailing = __builtin_ctz(n);
			n = n >> trailing;
			c += trailing-1;
		}

		c++;
	}

	return c;
}
