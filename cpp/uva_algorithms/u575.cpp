#include <iostream>
#include <algorithm>

#include <stdlib.h>

#include <vector>
#include <string>
#include <cmath>


using namespace std;


/**
 * UVA "skew binary" exercise <http://uva.onlinejudge.org/external/5/575.pdf>
 *
 * Convert input in skew binary to decimal.
 * each position represents the value 2^k+1
 *
 */
int main(int argc, const char * argv[])
{
    string skew_input;
    long skew_out;

    while (cin >> skew_input) {
    	if (skew_input == "0")
    		break;

    	// each position is n_k*(2^k)-1 for positions 0..k back to front
    	skew_out = 0;

    	// reverse iteration, calculate the position using std::distance (caveat: O(n) on non-random-access)
    	for(std::string::reverse_iterator it=skew_input.rbegin();  it < skew_input.rend(); it++) {
    		long i = std::distance(skew_input.rbegin(), it);

    		int p = pow(2, i+1) - 1;
    		// take the element
    		long pos = *it - '0';

    		skew_out += p*pos;
    	}

    	cout << skew_out << endl;
    }

    return 0;
}
