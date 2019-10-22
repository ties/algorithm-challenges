#include <iostream>

using namespace std;


/**
 * UVA "back to high school physics" <http://uva.onlinejudge.org/external/100/10071.pdf>
 *
 * input lines contain (v, t) pairs
 * the object has constant acceleration
 * what is the velocity after 2*t?
 *
 * constant acceleration
 * v = t*a
 * v_2 = 2*t*a
 *
 * (no need to derive the laws of motion here, whoops)
 */
int main(int argc, const char * argv[])
{
    long t1, v1;

    while (cin >> v1 >> t1) {
    	cout << 2*v1*t1 << endl;
    }
    return 0;
}
